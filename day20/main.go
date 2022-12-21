package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Elem struct {
	value      int
	elemP      *Elem
	prev, next *Elem
}

type DoubleLinkedList struct {
	head *Elem
	tail *Elem
	len  int
}

func (ddList *DoubleLinkedList) moveRight(elem *Elem) {
	//move elem right one step
	elemNext := elem.next
	elemPrev := elem.prev
	elemNextNext := elemNext.next

	elemNextNext.prev = elem
	elem.prev = elemNext
	elemNext.prev = elemPrev

	elemPrev.next = elemNext
	elemNext.next = elem
	elem.next = elemNextNext
	if ddList.head == elem {
		ddList.head = elemNext
	}
}

func (ddList *DoubleLinkedList) moveLeft(elem *Elem) {
	//move elem left one step
	elemNext := elem.next
	elemPrev := elem.prev
	elemPrevPrev := elemPrev.prev

	elemNext.prev = elemPrev
	elemPrev.prev = elem
	elem.prev = elemPrevPrev

	elemPrevPrev.next = elem
	elem.next = elemPrev
	elemPrev.next = elemNext
	if ddList.head == elem {
		ddList.tail = elem
		ddList.head = elemNext
	} else if ddList.head == elemPrev {
		ddList.tail = elem
	}
}

func (ddList *DoubleLinkedList) addHead(elem *Elem) {
	if ddList.head == nil {
		ddList.head = elem
		ddList.tail = elem
	} else {
		elem.next = ddList.head
		ddList.head.prev = elem
		ddList.head = elem
	}
	ddList.len++
}

func (ddList *DoubleLinkedList) addTail(elem *Elem) {
	if ddList.head == nil {
		ddList.head = elem
		ddList.tail = elem
	} else {
		// elem.next should be nil
		ddList.tail.next = elem
		elem.prev = ddList.tail
		ddList.tail = elem
	}
	ddList.len++
}

func (ddList *DoubleLinkedList) print() {
	dlPointer := ddList.head
	for i := 0; i < ddList.len; i++ {
		fmt.Printf("%v ", dlPointer.value)
		dlPointer = dlPointer.next
	}
}

func (ddList *DoubleLinkedList) printRev() {
	dlPointer := ddList.tail
	for i := 0; i < ddList.len; i++ {
		fmt.Printf("%v ", dlPointer.value)
		dlPointer = dlPointer.prev
	}
}
func main() {
	f, err := os.Open("testinput")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var zeroElem *Elem
	workDdList := new(DoubleLinkedList)
	iterDdList := new(DoubleLinkedList)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		elemInt, _ := strconv.Atoi(line)
		elem := Elem{
			value: elemInt,
			next:  nil,
			prev:  nil,
		}
		iterElem := Elem{
			elemP: &elem,
			value: elemInt,
			next:  nil,
			prev:  nil,
		}
		if elemInt == 0 {
			zeroElem = &elem
		}
		workDdList.addTail(&elem)
		iterDdList.addTail(&iterElem)
	}
	//fmt.Printf("WorkDdList: ")
	//workDdList.print()
	//fmt.Printf("\nIterDdList: ")
	//iterDdList.print()
	//fmt.Printf("\n")

	// Make the workDdList circular
	workDdList.head.prev = workDdList.tail
	workDdList.tail.next = workDdList.head

	nextIterElem := iterDdList.head
	for nextIterElem != nil {
		elem := nextIterElem.elemP
		if elem.value > 0 {
			for i := 0; i < elem.value; i++ {
				workDdList.moveRight(elem)
				if elem.value == 4 {
					workDdList.print()
					fmt.Printf("\n")
				}
			}
		} else if elem.value < 0 {
			valFloat := math.Abs(float64(elem.value))
			for i := 0; i < int(valFloat); i++ {
				workDdList.moveLeft(elem)
			}
		}
		//fmt.Printf("elem.value %v  moved ", elem.value)
		// workDdList.print()
		//fmt.Printf("   head %v tail %v\n", workDdList.head.value, workDdList.tail.value)
		nextIterElem = nextIterElem.next
	}
	//	fmt.Printf("Final list \n")
	// workDdList.print()

	var iterElem *Elem
	iterElem = zeroElem
	grooveCoords := 0
	for i := 1; i <= 3000; i++ {
		//fmt.Printf("%v %v ", i%1000, (i%1000) == 0)
		iterElem = iterElem.next
		if (i % 1000) == 0 {
			fmt.Printf("FFFFFFFF %v \n", iterElem.value)
			grooveCoords += iterElem.value
		}
	}
	fmt.Printf("%v \n", grooveCoords)
}
