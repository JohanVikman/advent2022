package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func printBoard(board [200][1000]string, deepestY int) {
	sands := 0
	for y := 0; y < deepestY; y++ {
		for x := 490; x < 585; x++ {
			if board[y][x] == "o" {
				sands++
			}
			fmt.Printf("%s", board[y][x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%v sands", sands)
}

func rockRange(board [200][1000]string, xstart int, ystart int, xend int, yend int) [200][1000]string {
	if xstart == xend {
		// column
		if yend < ystart {
			// up
			for i := 0; i < ystart-yend; i++ {
				board[ystart-i][xstart] = "#"
			}
		} else {
			for i := 0; i < yend-ystart; i++ {
				board[ystart+i][xstart] = "#"
			}
		}
	} else {
		// row
		if xend < xstart {
			for i := 0; i < xstart-xend; i++ {
				fmt.Printf("Adding row %v, %v\n", xstart-i, ystart)
				board[ystart][xstart-i] = "#"
			}
		} else {
			for i := 0; i < xend-xstart; i++ {
				fmt.Printf("Adding row %v, %v\n", xstart-i, ystart)
				board[ystart][xstart+i] = "#"
			}
		}
	}
	return board
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var board [200][1000]string
	for y := 0; y < 200; y++ {
		for x := 0; x < 1000; x++ {
			board[y][x] = "."
		}
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	// Build the board
	board[0][500] = "+"
	deepestY := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " -> ")
		first := true
		xprev := 0
		yprev := 0
		//fmt.Printf("%v %v\n", len(parts), parts)
		for i := 0; i < len(parts); i++ {
			part := parts[i]
			xy := strings.Split(part, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			if y > deepestY {
				deepestY = y
			}
			// fmt.Printf("%v, %v\n", x, y)
			if !first {
				//fmt.Printf("Add %v,%v -> %v, %v to board...\n", xprev, yprev, x, y)
				board = rockRange(board, xprev, yprev, x, y)
			}
			first = false
			xprev = x
			yprev = y
		}
	}
	//printBoard(board, 26)

	freeFalling := false
	sandUnits := 1
	//fmt.Printf("DeepestY = %v\n", deepestY)
	for {
		//sand := Point{500, 0}
		x := 500
		for y := 0; y <= deepestY; y++ {
			//fmt.Printf("x = %v y = %v ", x, y)
			if y == deepestY {
				freeFalling = true
				break
			}
			//fmt.Printf("%v,%v -> %v\n", x, y, board[y][x] == "#")
			if board[y+1][x] != "." {
				// should we stop or roll to left?
				if board[y+1][x-1] == "." && board[y][x-1] == "." {
					//roll left
					x--
					//fmt.Printf("continue...\n")
					continue
				} else if board[y+1][x+1] == "." && board[y][x+1] == "." {
					// roll right
					x++
					//fmt.Printf("continue...\n")
					continue
				} else {
					// stop
					sandUnits++
					board[y][x] = "o"
					break
				}
			}
		}

		if freeFalling {
			fmt.Printf("We are freeFalling!\n")
			break
		}
	}
	printBoard(board, deepestY+1)
	fmt.Printf("FreeFalling %v after %v\n", freeFalling, sandUnits)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
