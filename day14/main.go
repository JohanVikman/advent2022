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
	Y, X int
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
			for i := 0; i <= ystart-yend; i++ {
				board[ystart-i][xstart] = "#"
			}
		} else {
			for i := 0; i <= yend-ystart; i++ {
				board[ystart+i][xstart] = "#"
			}
		}
	} else if ystart == yend {
		// row
		if xend < xstart {
			for i := 0; i <= xstart-xend; i++ {
				board[ystart][xstart-i] = "#"
			}
		} else {
			for i := 0; i <= xend-xstart; i++ {
				board[ystart][xstart+i] = "#"
			}
		}
	} else {
		//ystart == yend and xstart ==  xend
		board[ystart][xstart] = "#"
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

	part1Sand := part1(board, deepestY)
	part2Sand := part2(board, deepestY+2)
	//printBoard(board, deepestY+1)
	fmt.Printf("FreeFalling after %v\n", part1Sand)
	fmt.Printf("Part2 %v\n", part2Sand)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func part1(board [200][1000]string, deepestY int) int {

	freeFalling := false
	sandUnits := 1
	fmt.Printf("DeepestY = %v\n", deepestY)
	for !freeFalling {
		sand := Point{0, 500}
		for {
			if sand.Y+1 > deepestY {
				freeFalling = true
				break
			}
			if board[sand.Y+1][sand.X] == "." {
				sand.Y++
			} else if board[sand.Y+1][sand.X-1] == "." {
				sand.X--
				sand.Y++
			} else if board[sand.Y+1][sand.X+1] == "." {
				sand.X++
				sand.Y++
			} else {
				board[sand.Y][sand.X] = "o"
				sandUnits++
				break
			}
		}

		if freeFalling {
			fmt.Printf("We are freeFalling!\n")
			break
		}
	}
	return sandUnits - 1
}

func part2(board [200][1000]string, floor int) int {

	for x := 0; x < 1000; x++ {
		board[floor][x] = "#"
	}

	printBoard(board, floor)
	sandUnits := 1
	for {
		sand := Point{0, 500}
		if board[sand.Y][sand.X] == "o" {
			//We've overwritten the "+"
			break
		}
		for {
			if board[sand.Y+1][sand.X] == "." {
				sand.Y++
			} else if board[sand.Y+1][sand.X-1] == "." {
				sand.X--
				sand.Y++
			} else if board[sand.Y+1][sand.X+1] == "." {
				sand.X++
				sand.Y++
			} else {
				board[sand.Y][sand.X] = "o"
				sandUnits++
				break
			}
		}
	}
	return sandUnits - 1
}
