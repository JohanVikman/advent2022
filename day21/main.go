package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Yell     = 1
	YellAdd  = 2
	YellSub  = 3
	YellProd = 4
	YellDiv  = 5
)

type Monkey struct {
	name string
	job  int
	val  int
	a    string
	b    string
}

func getMonkeyJob(parts []string) int {
	if len(parts) == 1 {
		return Yell
	} else {
		switch parts[1] {
		case "+":
			return YellAdd
		case "-":
			return YellSub
		case "*":
			return YellProd
		case "/":
			return YellDiv
		}
	}
	// for the compiler...
	return Yell
}

func (monkey Monkey) performOp(monkeyMap map[string]Monkey) int {
	var ret int
	if monkey.job == Yell {
		return monkey.val
	} else {
		monkeyA := monkeyMap[monkey.a]
		monkeyB := monkeyMap[monkey.b]
		switch monkey.job {
		case YellAdd:
			ret = monkeyA.val + monkeyB.val
		case YellSub:
			ret = monkeyA.val - monkeyB.val
		case YellProd:
			ret = monkeyA.val * monkeyB.val
		case YellDiv:
			ret = monkeyA.val / monkeyB.val
		}
	}
	monkey.val = ret
	return ret
}

func main() {
	f, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	monkeyMap := make(map[string]Monkey)
	rootAName := ""
	rootBName := ""

	//TODO: Algrebraic expansion.
	// e.g. we have root : aaaa + zzzz
	// we then get aaaa: bbbb - cccc
	// we can then exchange all aaaa for bbbb - cccc
	// when we can deduce a result for cccc, say cccc = 10
	// we can exchange bbbb - cccc = bbbb - 10
	// when we have a result for bbbb, say 20: bbbb-cccc=bbbb-10=20-10=10

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		name := strings.Trim(parts[0], ":")
		monkeyJob := getMonkeyJob(parts[1:])
		monkey := Monkey{name: name,
			job: monkeyJob}

		if monkeyJob == Yell {
			num, _ := strconv.Atoi(parts[1])
			monkey.val = num
			monkeyMap[name] = monkey
			fmt.Printf("Yell monkey %v\n", monkey)

		} else {
			aName := parts[1]
			bName := parts[3]
			monkey.job = monkeyJob
			monkey.a = aName
			monkey.b = bName
			monkey.performOp(monkeyMap)
			monkeyMap[monkey.name] = monkey
		}

		if name == "root" {
			rootAName = parts[1]
			rootBName = parts[3]
			fmt.Printf("found root! %v %v\n", rootAName, rootBName)
			//TODO: Check if we can perform the monkeyOp
			_, existsA := monkeyMap[rootAName]
			if existsA {
				_, existsB := monkeyMap[rootBName]
				if existsB {
					res := monkey.performOp(monkeyMap)
					fmt.Printf("Root %v\n", res)
					break
				}
			}
		}

		if rootAName != "" {
			if name == rootAName {
				fmt.Printf("rootAName found %v -> %v\n", monkey.name, monkey.val)
			}
		}
		if rootBName != "" {
			if name == rootBName {
				fmt.Printf("rootBName found %v -> %v\n", monkey.name, monkey.val)
			}
		}
		if rootAName != "" && rootBName != "" {
			rootMonkey, exists := monkeyMap["root"]
			if exists {
				rootMonkey.performOp(monkeyMap)
			}
		}
	}
}
