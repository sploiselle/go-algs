package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var n int

type Vertex struct {
	lat  float64
	long float64
}

func (v Vertex) String() string {
	return fmt.Sprintf("\nlat:\t%f\nlong:\t%f\n\n", v.lat, v.long)
}

var VertexMap = make(map[int]*Vertex)
var RemainingVertices = make(map[int]bool)

type edgeFromVertex struct {
	ID   int
	cost float64
}

func main() {

	readFile(os.Args[1])

	var tourCost float64

	currentEdge := 1
	delete(RemainingVertices, 1)
	chepestEdge := edgeFromVertex{0, math.Inf(1)}

	for len(RemainingVertices) > 1 {

		// reset chepestEdge.cost
		chepestEdge.cost = math.Inf(1)

		delete(RemainingVertices, currentEdge)

		for k := range RemainingVertices {

			thisDistance := distance(VertexMap[currentEdge], VertexMap[k])

			if thisDistance < chepestEdge.cost {
				chepestEdge.ID = k
				chepestEdge.cost = thisDistance
			} else if thisDistance == chepestEdge.cost {
				// problem solution is defined as using
				// vertex with lowest ID possible
				// and map iteration is random
				if k < chepestEdge.ID {
					chepestEdge.ID = k
					chepestEdge.cost = thisDistance
				}
			}
		}

		tourCost += chepestEdge.cost

		if chepestEdge.ID == currentEdge {
			break
		} else {
			currentEdge = chepestEdge.ID
		}

	}

	fmt.Println(int(tourCost + distance(VertexMap[1], VertexMap[chepestEdge.ID])))
}

func visitVertex(i int) {

}

func distance(a, b *Vertex) float64 {
	return math.Sqrt(math.Pow(a.lat-b.lat, 2) + math.Pow(a.long-b.long, 2))
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
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("couldn't convert number: %v\n", err)
		}
	}

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		thisID, err := strconv.Atoi(thisLine[0])
		thisLat, err := strconv.ParseFloat(thisLine[1], 64)
		thisLong, err := strconv.ParseFloat(thisLine[2], 64)

		if err != nil {
			log.Fatal(err)
		}

		VertexMap[thisID] = &Vertex{thisLat, thisLong}
		RemainingVertices[thisID] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
