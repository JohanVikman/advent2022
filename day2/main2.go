package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	//A, X - Rock,
	//B, Y - Paper,
	//C, Z - Scissors
	winningDict := make(map[string]string)
	winningDict["A"] = "Y"
	winningDict["B"] = "Z"
	winningDict["C"] = "X"

	drawDict := make(map[string]string)
	drawDict["A"] = "X"
	drawDict["B"] = "Y"
	drawDict["C"] = "Z"

	looseDict := make(map[string]string)
	looseDict["A"] = "Z"
	looseDict["B"] = "X"
	looseDict["C"] = "Y"

	pointDict := make(map[string]int)
	pointDict["X"] = 1
	pointDict["Y"] = 2
	pointDict["Z"] = 3

	strategyGuide := make(map[string]string)
	var sum = 0
	var index = 0
	var action string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		strategyGuide[parts[0]] = parts[1]
		fmt.Printf("%v -> %v \n", parts[0], parts[1])
		//Win or loose?
		if parts[1] == "X" {
			//loose
			fmt.Printf("Loose, choose %v\n", looseDict[parts[0]])
			action = looseDict[parts[0]]
		} else if parts[1] == "Y" {
			//draw
			fmt.Printf("draw, choose %v\n", looseDict[parts[0]])
			action = drawDict[parts[0]]
			sum += 3
		} else {
			//win
			fmt.Printf("win, choose %v\n", looseDict[parts[0]])
			action = winningDict[parts[0]]
			sum += 6
		}
		sum += pointDict[action]
		fmt.Printf("sum = %v\n", sum)
		index += 1
	}

	//6 win
	//3 draw
	//0 loss

	fmt.Printf("index: %v sum: %v", index, sum)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
