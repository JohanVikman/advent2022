package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func compareRanges(afrom int, bfrom int, ato int, bto int) bool {
	if afrom <= bfrom {
		if ato >= bto {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	dirSizes := make(map[string]uint32)
	var curDirSize uint32
	var totalSize uint32
	var pathLen int
	curDir := "/"
	var pathTokens []string
	var haveReadCurDir bool
	fmt.Printf("%v \n", haveReadCurDir)
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	line := fileScanner.Text()
	doReadNewLine := false
	for {
		if doReadNewLine == true {
			if fileScanner.Scan() == false {
				break
			}
			line = fileScanner.Text()
		} else {
			//Ok, but read a line in the next loop
			doReadNewLine = true
		}
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if len(parts) == 3 {
				path := parts[2]
				//Add to dirSizes
				if path == curDir {
					// We haven't moved, loop
					fmt.Printf("loop cd: %s %s from %s\n", parts[1], parts[2], curDir)
					continue
				} else if path == "/" {
					// Going to the top, save size and reset curDir
					fmt.Printf("reset cd: %s %s from %s\n", parts[1], parts[2], curDir)
					curDir = "/"
				} else if path == ".." {
					fmt.Printf("UP1 cd: %s %s from %s\n", parts[1], parts[2], curDir)
					//go up one level
					// TODO: Maybe handle cd .. from "/"
					dirSizes[curDir] = curDirSize
					pathTokens = strings.Split(curDir, "/")
					pathLen = len(pathTokens)
					curDir = strings.Join(pathTokens[:pathLen-1], "/") // '/' is not included
					_, exists := dirSizes[curDir]
					if exists {
						haveReadCurDir = true
					} else {
						haveReadCurDir = false
					}
				} else {
					//subdir
					fmt.Printf("DOWN1 cd: %s %s from %s\n", parts[1], parts[2], curDir)
					dirSizes[curDir] = curDirSize
					curDir += "/" + path
					fmt.Printf("New curDir = %s\n", curDir)
					_, exists := dirSizes[curDir]
					if exists {
						haveReadCurDir = true
					} else {
						haveReadCurDir = false
					}
				}
			} else if len(parts) == 2 {
				fmt.Printf("ls: \n")
				//read until we see a $
				fileScanner.Scan()
				line = fileScanner.Text()
				for {
					parts = strings.Split(line, " ")
					if parts[0] == "$" {
						doReadNewLine = false
						haveReadCurDir = true
						break
					} else {
						size, err := strconv.Atoi(parts[0])
						if err == nil {
							fmt.Printf("%s %v \n", parts[1], size)
							curDirSize += uint32(size)
						} else {
							fmt.Printf("%s is a dir? %T\n", parts[1], size)
						}
						if fileScanner.Scan() == false {
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

	//TODO: Go through all directories, find the big ones and add to totalSize
	fmt.Printf("totalSize = %v", totalSize)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
