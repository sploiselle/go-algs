package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numOfRows int

type Char struct {
	Weight   int
	ZeroSide *Char
	OneSide  *Char
}

func (c Char) String() string {

	return fmt.Sprintf("\n%v", c.Weight)
}

type CharHeap []*Char

func (ch CharHeap) Len() int { return len(ch) }

func (ch CharHeap) Less(i, j int) bool {
	return ch[i].Weight < ch[j].Weight
}

func (ch CharHeap) Swap(i, j int) {
	ch[i], ch[j] = ch[j], ch[i]
}

// Push adds Chares to CharHeaps
func (ch *CharHeap) Push(x interface{}) {
	*ch = append(*ch, x.(*Char))
}

// Pop returns the Char with the lowest DGS and removes it from the heap
func (ch *CharHeap) Pop() interface{} {
	old := *ch
	n := len(old)
	v := old[n-1]
	*ch = old[0 : n-1]
	return v
}

// type Tree struct {
// 	ZeroTree *Tree
// 	OneTree  *Tree
// 	ZeroChar *Char
// 	OneSide  *Char
// }

var ch CharHeap

var encodings []string

func main() {

	heap.Init(&ch)

	readFile(os.Args[1])

	finalTree := compress(&ch)

	printTraversal("\n", finalTree)

	encodingMin := len(encodings[0])
	encodingMax := 0

	for _, v := range encodings {

		lv := len(v) - 1

		// fmt.Printf("%s\t%d", v, lv)

		if lv < encodingMin {
			encodingMin = lv
		}

		if lv > encodingMax {
			encodingMax = lv
		}
	}

	fmt.Printf("%d\n%d", encodingMax, encodingMin)
}

func readFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan first line
	if scanner.Scan() {
		numOfRows, err = strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatalf("couldn't convert number: %v\n", err)
		}
	}

	for scanner.Scan() {

		thisWeight, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		c := &Char{thisWeight, nil, nil}

		heap.Push(&ch, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compress(ch *CharHeap) *Char {

	a := heap.Pop(ch).(*Char)
	b := heap.Pop(ch).(*Char)

	if len(*ch) == 0 {
		return &Char{(a.Weight + b.Weight), a, b}
	}

	heap.Push(ch, &Char{(a.Weight + b.Weight), a, b})

	return compress(ch)
}

func printTraversal(encodingSoFar string, c *Char) {

	if c.OneSide == nil {
		encodings = append(encodings, encodingSoFar)
	} else {
		printTraversal(encodingSoFar+"1", c.OneSide)
	}

	if c.ZeroSide == nil {
		encodings = append(encodings, encodingSoFar)
	} else {
		printTraversal(encodingSoFar+"0", c.ZeroSide)
	}
}
