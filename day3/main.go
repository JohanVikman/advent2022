package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func addToMap(i int32, mymap map[int32]int32) {
	val, exists := mymap[i]
	if exists {
		mymap[i] = val + 1
	} else {
		mymap[i] = 1
	}
}

func convertChar(c int32) int32 {
	var result int32
	if c >= 97 {
		result = c - 96
	} else {
		result = c - 38
	}
	return result
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var sum int32
	var length int
	var lineno int32

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		length = len(line)
		var firstHalf = make(map[int32]int32)
		var secondHalf = make(map[int32]int32)
		fmt.Printf("line \"%v\" (%d)(%d)\n", line, len(line), lineno)

		for ix, item := range line {
			if ix < length/2 {
				fmt.Printf("Adding to 1st half %v: %s\n", item, string(item))
				addToMap(item, firstHalf)
			} else {
				fmt.Printf("Adding to 2nd half %v: %s\n", item, string(item))
				addToMap(item, secondHalf)
			}
		}

		for key, _ := range firstHalf {
			_, exists := secondHalf[key]
			if exists {
				fmt.Printf("%c exists in both, adding %d\n", key, convertChar(key))
				sum += convertChar(key)
			}
		}
		fmt.Printf("len(firstHalf) = %v, len(secondHalf) = %v\n\n", len(firstHalf), len(secondHalf))
		lineno += 1
	}

	fmt.Printf("sum = %v", sum)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
