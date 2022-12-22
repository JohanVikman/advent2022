package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func compareRanges(afrom int, bfrom int, ato int, bto int) bool {
	if afrom <= bfrom {
		if ato >= bto {
			return true
		}
	}
	return false
}

//rotate right
func matrixRotate(matrix [5][5]int) [5][5]int {
	var target [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			target[j][5-i-1] = matrix[i][j]
		}
	}
	return target
}

func main() {
	f, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// 99 rows.
	// 99
	size := 5
	var forest [5][5]int
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for i := 0; i < size; i++ {
		fileScanner.Scan()
		line := fileScanner.Text()
		for j := 0; j < size; j++ {
			tree, err := strconv.Atoi(string(line[j]))
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			forest[i][j] = tree
		}
	}
	// TODO: Check forest from each direction by rotating the matrix
	seenHeight := 0
	seen := 0
	fmt.Printf("Forest = %v\n", forest)
	for rot := 0; rot < 4; rot++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if forest[i][j] < 0 {
					fmt.Printf("X")
					if -(forest[i][j]) > seenHeight {
						seenHeight = -(forest[i][j])
						continue
					} else {
						continue
					}
				}
				if seenHeight < forest[i][j] {
					seen += 1
					seenHeight = forest[i][j]
					forest[i][j] = -seenHeight
					fmt.Printf("1")
				} else if seenHeight == 0 && forest[i][j] == 0 {
					seen += 1
					forest[i][j] = -seenHeight
					fmt.Printf("1")

				} else {
					fmt.Printf("0")
				}
			}
			fmt.Printf("\n")

			seenHeight = 0
		}
		seenHeight = 0
		fmt.Printf("Forest = %v\n", forest)
		forest = matrixRotate(forest)
		fmt.Printf("rotated = %v\n", matrixRotate(forest))
	}
	//TODO: Go through all directories, find the big ones and add to totalSize
	fmt.Printf("seen = %v\n", seen)
	//fmt.Printf("rotated = %v\n", matrixRotate(forest))
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
