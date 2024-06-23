package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// var inputMap = make(map[int]bool)

var inputArray []int

var tMap = make(map[int]bool)

//IotaError identifies numbers that failed to be convered
type IotaError struct {
	Num string
}

//Error stringer for IotaError
func (e *IotaError) Error() string {
	return fmt.Sprintf("Couldn't convert %v", e.Num)
}

func main() {

	err := readFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(twoSums(-10000, 10000))
}

func readFile(filename string) error {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		thisInt, err := strconv.Atoi(thisLine[0])

		if err != nil {
			return &IotaError{thisLine[0]}
		}

		inputArray = append(inputArray, thisInt)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(inputArray)
	return nil
}

func twoSums(lowerT int, upperT int) int {

	for _, x := range inputArray {

		lowerTIndex := sort.Search(len(inputArray), func(i int) bool { return inputArray[i] >= (lowerT - x) })

		upperTIndex := sort.Search(len(inputArray), func(i int) bool { return inputArray[i] > (upperT - x) })

		for i := lowerTIndex; i < upperTIndex; i++ {

			thisT := x + inputArray[i]

			// fmt.Printf("Found a T:\t%d", thisT)

			_, ok := tMap[thisT]

			if !ok {
				tMap[thisT] = true
			}
		}
	}

	return len(tMap)
}
