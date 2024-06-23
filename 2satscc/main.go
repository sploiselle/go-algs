package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	Id       int
	Edges    []int
	Explored bool
}

func (v Vertex) String() string {
	return fmt.Sprintf("id:\t%d\nedges: %v\n\n", v.Id, v.Edges)
}

func (v *Vertex) AddEdge(i int) {
	v.Edges = append(v.Edges, i)
}

type Leader struct {
	Id      int
	Members map[int]bool
}

func (v Leader) String() string {
	return fmt.Sprintf("id:\t%d\nmembers: %v\n\n", v.Id, v.Members)
}

func (v *Leader) AddMember(i int) {
	v.Members[i] = true
}

// Adjacency lists
var vertexMap_f = make(map[int]*Vertex)
var vertexMap_b = make(map[int]*Vertex)

// Ensures all nodes are hit in first passthrough
var backwardKeys = make(map[int]int)

// The Secret Sauce™
var magicalOrderMap = make(map[int]int)

// Tracks SCC leaders
var leaderMap = make(map[int]*Leader)

// Pointer to current leader
var s *Leader

// Tracks funning time
var t int

func main() {

	readFile(os.Args[1])

	DFSLoop(vertexMap_b, backwardKeys, 1)
	DFSLoop(vertexMap_f, magicalOrderMap, 2)

	// fmt.Print("\n\n\n LEADERS\n\n%v", leaderMap)

	for _, v := range leaderMap {
		for w := range v.Members {
			_, ok := v.Members[w*-1]

			if ok {
				fmt.Println("0")
				os.Exit(1)
			}
		}
	}

	fmt.Println("1")

	// var leaderOrder []int

	// for _, v := range leaderMap {
	// 	leaderOrder = append(leaderOrder, len(v.Members))
	// }

	// sort.Sort(sort.Reverse(sort.IntSlice(leaderOrder)))

	// fmt.Print("\n%v\n", leaderOrder[0:5])
}

func readFile(filename string) {

	//Used as the 'key' in backwardsKeys map
	k := 0

	file, err := os.Open(filename) //should read in file named in CLI web ew we we we we  we w ew e wwe wewe
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan first line
	if scanner.Scan() {
		numOfRows, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatalf("couldn't convert number: %v\n", err)
		}

		for i := 1; i <= numOfRows; i++ {

			ni := i * -1
			vf := &Vertex{i, []int{}, false}
			vertexMap_f[i] = vf

			nvf := &Vertex{ni, []int{}, false}
			vertexMap_f[ni] = nvf

			// fmt.Println(vertexMap_f[ni])

			vb := &Vertex{i, []int{}, false}
			vertexMap_b[i] = vb

			nvb := &Vertex{ni, []int{}, false}
			vertexMap_b[ni] = nvb

			backwardKeys[k] = i
			k++
			backwardKeys[k] = ni
			k++
		}
	}

	for scanner.Scan() {

		thisLine := strings.Fields(scanner.Text())

		sat1, err := strconv.Atoi(thisLine[0])
		sat2, err := strconv.Atoi(thisLine[1])

		if err != nil {
			fmt.Print("couldn't convert number: %v\n", err)
			return
		}

		nsat1V, ok := vertexMap_f[sat1*-1]

		if !ok {
			log.Fatal("Couldn't find ", sat1*-1)
		}
		nsat2V, ok := vertexMap_f[sat2*-1]

		if !ok {
			log.Fatal("Couldn't find ", sat1*-1)
		}

		nsat1V.AddEdge(sat2)
		nsat2V.AddEdge(sat1)

		sat1V, ok := vertexMap_b[sat1]

		if !ok {
			log.Fatal("Couldn't find ", sat1)
		}

		sat2V, ok := vertexMap_b[sat2]

		if !ok {
			log.Fatal("Couldn't find ", sat2)
		}

		sat1V.AddEdge(sat2 * -1)
		sat2V.AddEdge(sat1 * -1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func DFSLoop(graph map[int]*Vertex, orderer map[int]int, pass int) {

	//fmt.Print("\nFor pass %d, orderer is %v", pass, orderer)

	for i := len(orderer); i > 0; i-- {

		//fmt.Print("\n\nFor pass %d looking for %d", pass, orderer[i])
		w, ok := graph[orderer[i]]

		if ok {

			if !w.Explored {

				if pass == 2 {
					// //fmt.Print("\n\n\nFound Leader %d", w.Id)
					s = &Leader{w.Id, make(map[int]bool)}
					leaderMap[s.Id] = s
				}

				DFS(graph, w, pass)
			}

		}
	}
}

func DFS(graph map[int]*Vertex, i *Vertex, pass int) {

	i.Explored = true

	if pass == 2 {
		s.AddMember(i.Id)
	}

	for _, v := range i.Edges {

		vertex := graph[v]

		if !vertex.Explored {
			DFS(graph, vertex, pass)
		}
	}

	if pass == 1 {
		t++
		magicalOrderMap[t] = i.Id
	}

}
