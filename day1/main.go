package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	max := 0
	cur := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			fmt.Printf("empty line cur %v max %v\n", cur, max)
			if cur > max {
				max = cur
			}
			cur = 0
			continue
		}
		v, _ := strconv.Atoi(line)
		cur += v
	}

	fmt.Printf("max %v", max)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
