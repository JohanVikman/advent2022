package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	file uint8 = iota
	dir
)

type Tree struct {
	tpe      uint8  `file or dir`
	tag      string `filename or dirname`
	size     int32  `size of file or size of all files and directories below this node`
	parent   *Tree
	children map[string]*Tree `all files and subdirectories`
}

func NewTree(str string, parent *Tree) *Tree {
	var t Tree
	t.tpe = dir
	t.tag = str
	t.size = 0
	t.parent = parent
	t.children = make(map[string]*Tree)
	return &t
}

func (tree *Tree) Unwind(sz int32) {
	fmt.Printf("Adding %v to %v\n", sz, tree.tag)
	tree.size += sz
	if tree.parent != nil {
		fmt.Printf("Added %v to %v\n", sz, tree.tag)
		tree.parent.Unwind(sz)
	} else {
		fmt.Printf("Reached top at %v\n", tree.tag)
	}
}

func (tree *Tree) Add(node Tree) {
	node2, exists := tree.children[node.tag]
	if exists {
		if node2.size != node.size && node.size != 0 {
			diff := node2.size - node.size
			tree.Unwind(diff)
		}
	} else {
		tree.children[node.tag] = &node
		tree.Unwind(node.size)
	}
}

func (tree *Tree) AddFrom(fromPath []string, path []string, node Tree) bool {
	node2, exists := tree.children[node.tag]
	if exists {
		fmt.Printf("Already added %v\n", node.tag)
		// TODO: is the size the same?
		if node2.size != node.size {
			diff := node2.size - node.size
			tree.Unwind(diff)
			return true
		} else {
			//Nothing to do
			return true
		}
	}
	if len(path) == 0 {
		node.parent = tree
		tree.children[node.tag] = &node
		tree.Unwind(node.size)
		return true
	} else {
		headPath := path[0]
		restPath := path[1:]
		fmt.Printf("headPath = %v\nrestPath = %v\n", headPath, restPath)
		childPath, exists := tree.children[path[0]]
		if exists {
			childPath.Add(node)
		} else {
			fmt.Printf("Path %s does not exist\n", headPath)
			return false
		}
	}
	return false
}

func (tree *Tree) Find(path []string) (bool, *Tree) {
	if len(path) == 0 {
		foundNode, exists := tree.children[path[0]]
		if exists {
			return true, foundNode
		} else {
			fmt.Printf("Could not find %s\n", path[0])
			return false, nil
		}
	} else {
		headPath := path[0]
		restPath := path[1:]
		childPath, exists := tree.children[path[0]]
		if exists {
			childPath.Find(restPath)
		} else {
			fmt.Printf("Path %s does not exist\n", headPath)
			return false, nil
		}
	}
	return false, nil
}

func (tree *Tree) Cd(path []string) (bool, *Tree) {
	if len(path) == 0 {
		fmt.Printf("Reached end...\n")
		return true, tree
	}

	fmt.Printf("Cd '%v' %v\n", path[0], len(path))

	if path[0] == "." && len(path) == 1 {
		return true, tree
	} else if path[0] == "/" && len(path) == 1 {
		if tree.parent == nil {
			return true, tree
		} else {
			fmt.Printf("tree parent %v\n", tree.parent)
			ok, ret := tree.parent.Cd(path)
			if ok {
				return true, ret
			} else {
				return false, nil
			}
		}
	} else if path[0] == "" {
		if tree.parent == nil {
			return true, tree
		} else {
			fmt.Printf("tree parent %v\n", tree.parent)
			ok, ret := tree.parent.Cd(path)
			if ok {
				return true, ret
			} else {
				return false, nil
			}
		}
	} else if path[0] == ".." {
		if tree.parent != nil {
			fmt.Printf("Trying %v\n", path[1:])
			ok, ret := tree.parent.Cd(path[1:])
			if ok {
				return true, ret
			} else {
				return false, nil
			}
		} else {
			fmt.Printf("cannot go to %v from root\n", path)
			return false, nil
		}
	} else {
		fmt.Printf("Looking for %v in %v\n", path[0], tree.children)
		childPath, exists := tree.children[path[0]]
		if exists {
			fmt.Printf("#1 %v (cd... %v)\n", childPath, path[1:])
			return childPath.Cd(path[1:])
		} else {
			fmt.Printf("#2 cannot go to %v from %v\n", path[0], tree.tag)
			return false, nil
		}
	}
}

func printWs(tab int) {
	for i := 0; i < tab; i++ {
		fmt.Print(" ")
	}
}

func (tree *Tree) Print(tab int) {
	printWs(tab)
	fmt.Printf("\\-%v %v \n", tree.tag, tree.size)
	for _, child := range tree.children {
		if child.tpe == dir {
			child.Print(tab + 2)
		} else {
			printWs(tab + 2)
			fmt.Printf("|-%v %v\n", child.tag, child.size)
		}
	}
}

func (tree *Tree) Walk(f func(t *Tree)) {
	for _, child := range tree.children {
		if child.tpe == dir {
			child.Walk(f)
		}
	}
	f(tree)
}

var totalSize int32

func sumUpSize(t *Tree) {
	fmt.Printf("Looking at %v %v\n", t.tag, t.size)
	if t.size <= 100000 {
		totalSize += t.size
	}
}

func compareRanges(afrom int, bfrom int, ato int, bto int) bool {
	if afrom <= bfrom {
		if ato >= bto {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	line := fileScanner.Text()
	doReadNewLine := false
	fileTree := NewTree("/", nil)
	for {
		if doReadNewLine {
			if !fileScanner.Scan() {
				break
			}
			line = fileScanner.Text()
			fmt.Printf("Read new line %v..\n", line)
		} else {
			//Ok, but read a line in the next loop
			fmt.Printf("Continue with the last line\n")
			doReadNewLine = true
		}
		parts := strings.Split(line, " ")
		fmt.Printf("====\n")
		fmt.Printf("Read \"%v\"\n", line)
		fileTree.Print(0)
		if parts[0] == "$" {
			if parts[1] == "cd" {
				pParts := strings.Split(parts[2], "/")
				for i := 0; i < len(pParts); i++ {
					tag := pParts[i]
					fmt.Printf("...%v (%v of %v)\n", tag, i+1, len(pParts))
					_, exists := fileTree.children[tag]
					if tag != "" && !exists {
						newNode := NewTree(tag, fileTree)
						fileTree.Add(*newNode)
						fmt.Printf("Found new subdir %v, adding..\n", tag)
						fileTree.Print(0)
						fmt.Printf("-----\n")
					}
					ok, newFileTree := fileTree.Cd(pParts)
					if ok {
						fileTree = newFileTree
					} else {
						fmt.Printf("ERROR CD\n")
						log.Fatal(-1)
					}
				}
				fmt.Printf("Cd done...\n")
			} else if parts[1] == "ls" {
				// read until we see a $
				fileScanner.Scan()
				line = fileScanner.Text()
				fmt.Printf("%v\n", line)
				for {
					parts = strings.Split(line, " ")
					if parts[0] == "$" {
						// We've reached the
						// end and there's a
						// new command in
						// 'line'. break the for-loop and start over
						doReadNewLine = false
						break
					} else {
						size, err := strconv.Atoi(parts[0])
						// if 'err' is nil, we've seen a directory
						isFile := err == nil
						tag := parts[1]
						_, exists := fileTree.children[tag]
						if isFile {
							if !exists {
								newNode := NewTree(tag, fileTree)
								newNode.tpe = file
								newNode.size = int32(size)
								fileTree.Add(*newNode)
								fmt.Printf("Added %s %v => tree.size = %v\n", parts[1], size, fileTree.size)
							}
						} else {
							if !exists {
								newNode := NewTree(tag, fileTree)
								fileTree.Add(*newNode)
							}
						}
						if !fileScanner.Scan() {
							// HACK
							doReadNewLine = true
							break
						}
						line = fileScanner.Text()
					}
				}
			}
		}
	}

	fmt.Printf("\n-=-=-=-=-=-=-=-=-=\nALL DONE GO THROUGH EVERYTHING\n")
	_, fileTree = fileTree.Cd([]string{"/"})
	fileTree.Print(0)
	fileTree.Walk(sumUpSize)

	//TODO: Go through all directories, find the big ones and add to totalSize
	fmt.Printf("totalSize = %v", totalSize)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
