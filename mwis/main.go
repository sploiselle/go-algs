package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var numOfRows int

var path []int
var A []int

func main() {

	readFile(os.Args[1])

	mwisPath := make(map[int]int, numOfRows)

	A[0] = 0
	A[1] = path[1]

	for i := 2; i < len(A); i++ {

		A[i] = int(math.Max(float64(A[i-1]), float64(A[i-2]+path[i])))

	}

	i := len(A) - 1

	// fmt.Println(i)
	// fmt.Println(len(path))

	for i >= 2 {

		// last := A[i-1]
		// this := A[i-2] + path[i]

		// fmt.Printf("\nthis\t%d\nlast\t%d\n", this, last)

		if A[i-1] >= A[i-2]+path[i] {
			i--
		} else {
			mwisPath[i] = A[i-2] + path[i]
			i -= 2
		}
	}

	if i == 1 {
		mwisPath[1] = path[i]
	}

	weirdResults := []int{1, 2, 3, 4, 17, 117, 517, 997}

	var results string

	for _, x := range weirdResults {

		_, ok := mwisPath[x]

		if !ok {
			results += "0"
		} else {
			results += "1"
		}
	}

	// fmt.Println(path)
	// fmt.Println(A)
	// fmt.Println(mwisPath)
	fmt.Println(results)
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

		path = make([]int, numOfRows+1)
		A = make([]int, numOfRows+1)
	}

	i := 1

	for scanner.Scan() {

		thisWeight, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		path[i] = thisWeight
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
