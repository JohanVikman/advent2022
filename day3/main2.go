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
	var lineno = 1

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var first = make(map[int32]int32)
	var second = make(map[int32]int32)
	var third = make(map[int32]int32)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		for _, item := range line {
			if lineno == 1 {
				addToMap(item, first)
			} else if lineno == 2 {
				addToMap(item, second)
			} else {
				addToMap(item, third)
			}
		}

		if lineno == 3 {
			for key := range first {
				_, exists := second[key]
				if exists {
					_, existsThird := third[key]
					if existsThird {
						sum += convertChar(key)
					}
				}
			}
			lineno = 1
			for k := range first {
				delete(first, k)
			}
			for k := range second {
				delete(second, k)
			}
			for k := range third {
				delete(third, k)
			}
		} else {
			lineno += 1
		}
	}

	fmt.Printf("sum = %v\n", sum)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
