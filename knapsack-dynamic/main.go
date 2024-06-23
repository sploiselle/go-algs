package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var W int // Knapsack weight capacity
var N int // Number of items in input

type Item struct {
	Value  int
	Weight int
}

var itemMap map[int]*Item

func main() {

	readFile(os.Args[1])

	//Include item 0/no items (N+1)
	a := make([][]int, N+1)

	for i := range a {
		//Include "max Value for 0 weight" (W+1)
		a[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {

		thisItem := itemMap[i]

		for j := 1; j <= W; j++ {

			wPrime := j - thisItem.Weight

			if wPrime < 0 {
				//Cannot choose item if its weight is greater
				//than the weight being considered
				a[i][j] = a[i-1][j]
			} else {
				// Choose index to the left
				// or best solution for wPrime
				// plus this item's weight
				a[i][j] = max(a[i-1][j], (a[i-1][wPrime] + thisItem.Value))
			}
		}
	}

	// Maximum value
	fmt.Println(a[N][W])
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

		firstLine := strings.Fields(scanner.Text())

		W, err = strconv.Atoi(firstLine[0])
		N, err = strconv.Atoi(firstLine[1])

		if err != nil {
			log.Fatalf("couldn't convert number: %v\n", err)
		}

		itemMap = make(map[int]*Item, N)

	}

	// Number for map, beginning with 1
	i := 1

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		val, err := strconv.Atoi(thisLine[0])
		weight, err := strconv.Atoi(thisLine[1])

		if err != nil {
			log.Fatal(err)
		}

		v, ok := itemMap[i]

		if ok {
			log.Fatal(v)
		} else {
			itemMap[i] = &Item{val, weight}
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func max(x, y int) int {

	if x > y {
		return x
	}
	return y
}
