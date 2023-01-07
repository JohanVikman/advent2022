package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Tree struct {
	height int
	seen   bool
}

func createTree(height int, seen bool) *Tree {
	t := Tree{height: height,
		seen: seen,
	}
	return &t
}

func compareRanges(afrom int, bfrom int, ato int, bto int) bool {
	if afrom <= bfrom {
		if ato >= bto {
			return true
		}
	}
	return false
}

// rotate right
func matrixRotate(matrix [99][99]*Tree) [99][99]*Tree {
	var target [99][99]*Tree
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			target[j][99-i-1] = matrix[i][j]
		}
	}
	return target
}
func printForest(forest [99][99]*Tree) {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			fmt.Printf("%v", forest[i][j].height)
		}
		fmt.Printf("\n")
	}

}

func printSeenTrees(forest [99][99]*Tree) {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			if forest[i][j].seen {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Printf("\n")
	}

}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// 99 rows.
	// 99
	size := 99
	var forest [99][99]*Tree
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for i := 0; i < size; i++ {
		fileScanner.Scan()
		line := fileScanner.Text()
		for j := 0; j < size; j++ {
			treeH, err := strconv.Atoi(string(line[j]))
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			tree := createTree(treeH, false)
			forest[i][j] = tree

		}
	}
	seenHeight := 0
	seen := 0
	// printForest(forest)
	fmt.Printf("READ OK\n\n\n")
	for rot := 0; rot < 4; rot++ {
		for i := 0; i < 99; i++ {
			for j := 0; j < 99; j++ {
				tree := forest[i][j]
				treeHeight := tree.height
				if i == 0 {
					// All trees on the edges can
					// be seen
					if !tree.seen {
						// Corner trees can be
						// seen from multiple
						// directions, only
						// count them once
						seen += 1
						tree.seen = true
					}
				} else if seenHeight < treeHeight {
					if !tree.seen {
						seen += 1
						tree.seen = true
					}
					seenHeight = treeHeight
				}
			}
			seenHeight = 0
		}
		seenHeight = 0
		forest = matrixRotate(forest)

	}
	//printSeenTrees(forest)
	fmt.Printf("\n")
	fmt.Printf("seen = %v\n", seen)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
