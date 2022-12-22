package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type GridVisit struct {
	grid    [4][4]bool
	visited int
}

func adjacent(hX int, hY int, tX int, tY int) bool {
	//determine if the two nodes are adjacent or not
	if math.Abs(float64(hX-hY)) <= 1 &&
		math.Abs(float64(hY-tY)) <= 1 {
		return true
	}
	return false
}

func markAsVisited(tX int, tY int, visited GridVisit) GridVisit {
	if visited.grid[tX][tY] != true {
		visited.visited++
		visited.grid[tX][tY] = true
		return visited
	}
	return visited
}

func main() {
	f, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	size := 4
	//var moveTail bool
	// store both visited true or false and the number of visited
	var visited GridVisit
	visited.grid[0][0] = true
	visited.visited = 0
	var hX int
	var hY int
	var tX int
	var tY int
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; i < size; i++ {
		fileScanner.Scan()
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		// Move H Up, down, left or right
		steps, _ := strconv.Atoi(parts[1])
		for j := 0; j < steps; j++ {
			if parts[0] == "U" {
				hY++
			} else if parts[0] == "D" {
				hY--
			} else if parts[0] == "L" {
				hY--
			} else if parts[0] == "R" {
				hY++
			}
			// Check if tail is adjacent, if not, move tail
			fmt.Printf("%v,%v and %v,%v adjacent -> %b\n", hX, hY, tX, tY, adjacent(hX, hY, tX, tY))
			if !adjacent(hX, hY, tX, tY) {
				if hX > tX {
					tX++
				} else if hX < tX {
					tX--
				}
				if hY > tY {
					tY++
				} else if hY < tY {
					tY--
				}
				visited = markAsVisited(tX, tY, visited)
			}
		}
	}
	fmt.Printf("Grid %v\n", grid)
	//fmt.Printf("rotated = %v\n", matrixRotate(forest))
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
