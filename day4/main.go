package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func compareRanges(afrom int, bfrom int, ato int, bto int) (bool) {
	if afrom <= bfrom {
		if ato >= bto {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var sum int32

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		allParts := strings.Split(line, ",")
		firstParts := strings.Split(allParts[0], "-")
		secondParts := strings.Split(allParts[1], "-")
		firstFrom, _ := strconv.Atoi(firstParts[0])
		firstTo, _ := strconv.Atoi(firstParts[1])
		secondFrom, _ := strconv.Atoi(secondParts[0])
		secondTo, _ := strconv.Atoi(secondParts[1])
		if (compareRanges(firstFrom, secondFrom, firstTo, secondTo) ||
			compareRanges(secondFrom, firstFrom, secondTo, firstTo)) {
			fmt.Printf("%v is one\n", line)
			sum += 1
		}
	}

	fmt.Printf("sum = %v", sum)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
