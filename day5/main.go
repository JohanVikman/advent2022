package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	//"io"
	"strconv"
)

type parseMode uint8

type Stack struct {
	items []string
	size  uint8
}

const (
	crate parseMode = 0
	move  parseMode = 1
)

func Push(st Stack, elem string) Stack {
	st.items = append([]string{elem}, st.items...)
	st.size++
	return st
}

func revPush(st Stack, elem string) Stack {
	st.items = append(st.items, elem)
	st.size++
	return st
}

func Pop(st Stack) (Stack, string) {
	if st.size == 0 {
		return st, "ö"
	} else {
		item := st.items[0]
		st.items = st.items[1:st.size]
		st.size--
		return st, item
	}
}

func parseStacks(cratesF *os.File, noStacks int, stacks []Stack) {
	cratesF.Seek(0, 0)
	fileScanner := bufio.NewScanner(cratesF)
	fileScanner.Split(bufio.ScanLines)
	maxLength := 3*noStacks + (noStacks - 1)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for ix := 0; ix < maxLength; ix += 4 {
			if line[ix+1] != 32 {
				crate := string(line[ix+1])
				stackIndex := ix / 4
				newStack := revPush(stacks[stackIndex], crate)
				stacks[stackIndex] = newStack
			}
		}
	}
}

func moveOp(stacks []Stack, numCrates int, fromStack int, toStack int) []Stack {
	for i := 0; i < numCrates; i++ {
		newFrom, popped := Pop(stacks[fromStack])
		newTo := Push(stacks[toStack], popped)
		stacks[fromStack] = newFrom
		stacks[toStack] = newTo
	}
	return stacks
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var parseMode = crate
	var stacks []Stack
	for i := 0; i < 9; i++ {
		stacks = append(stacks, Stack{})
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	cratesF, err := os.Create("crates")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer cratesF.Close()

	regexp, err := regexp.Compile(`\d`)

	if err != nil {
		log.Fatal(err)
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			//skip the empty line, we change parseMode when we find the crate stack indexes
			continue
		}
		if parseMode == crate {
			if regexp.MatchString(line) {
				//TODO: Read the saved stacks and enter them into the stacs array
				parseStacks(cratesF, 9, stacks)
				parseMode = move
				continue
			}
			_, err := cratesF.Write([]byte(line + "\n"))
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

		} else {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			fromIx, _ := strconv.Atoi(parts[3])
			toIx, _ := strconv.Atoi(parts[5])
			fromIx--
			toIx--
			stacks = moveOp(stacks, num, fromIx, toIx)
		}
	}
	fmt.Printf("%s", stacks[0].items[0])
	fmt.Printf("%s", stacks[1].items[0])
	fmt.Printf("%s", stacks[2].items[0])
	fmt.Printf("%s", stacks[3].items[0])
	fmt.Printf("%s", stacks[4].items[0])
	fmt.Printf("%s", stacks[5].items[0])
	fmt.Printf("%s", stacks[6].items[0])
	fmt.Printf("%s", stacks[7].items[0])
	fmt.Printf("%s\n", stacks[8].items[0])
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
