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

	max1 := 0
	max1prev := 0
	max2 := 0
	max2prev := 0
	max3 := 0
	cur := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			fmt.Printf("empty line cur %v max %v %v %v\n", cur, max1, max2, max3)
			if cur > max1 {
				max1 = cur
				max2 = max1prev
				max3 = max2prev
				max1prev = max1
				max2prev = max2
			} else if cur > max2 {
				max2 = cur
				max3 = max2prev
				max2prev = max2
			} else if cur > max3 {
				max3 = cur
			}
			cur = 0
			continue
		}
		v, _ := strconv.Atoi(line)
		cur += v
	}

	fmt.Printf("max %v", max1+max2+max3)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
