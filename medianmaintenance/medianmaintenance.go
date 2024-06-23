//Package medianmaintenance implements the 'median maintenance' algorithm for a list of numbers.
//This particular implementation uselessly sums all of the medians from the stream
//and then %1000
package medianmaintenance

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//maxHeap maintains the lesser half of the elements
type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h maxHeap) Peek() interface{} {
	n := len(h)

	if n == 0 {
		return 0
	}

	x := int(h[0])
	return x
}

//minHeap maintains the greater half of the elements
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Peek() interface{} {

	n := len(h)

	if n == 0 {
		return 0
	}

	x := int(h[0])

	return x
}

//IotaError identifies numbers that failed to be convered
type IotaError struct {
	Num string
}

//Error stringer for IotaError
func (e *IotaError) Error() string {
	return fmt.Sprintf("Couldn't convert %v", e.Num)
}

var lowerHalf = new(maxHeap)
var upperHalf = new(minHeap)

var results = make([]int, 0)

//GenerateFromFile creates median maintenance for values from a file
//where each line represents another int
func GenerateFromFile(filename string) (int, error) {

	heap.Init(lowerHalf)
	heap.Init(upperHalf)

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
			return 0, &IotaError{thisLine[0]}
		}

		medianMaintenance(thisInt, lowerHalf, upperHalf, &results)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return computeAnswer(results), nil
}

//GenerateFromIntSlice creates median maintenance for values from an
//[]int where each element represents another int
func GenerateFromIntSlice(s []int) (int, error) {

	heap.Init(lowerHalf)
	heap.Init(upperHalf)

	for _, v := range s {
		medianMaintenance(v, lowerHalf, upperHalf, &results)
	}

	return computeAnswer(results), nil
}

func medianMaintenance(i int, b *maxHeap, t *minHeap, res *[]int) {

	if i > b.Peek().(int) {
		heap.Push(t, i)
	} else {
		heap.Push(b, i)
	}

	rebalance(b, t)

	*res = append(*res, b.Peek().(int))
}

func rebalance(b *maxHeap, t *minHeap) {

	if b.Len() < t.Len() {
		heap.Push(b, heap.Pop(t))
	}

	if (t.Len() + 1) < b.Len() {
		heap.Push(t, heap.Pop(b))
	}

}

func computeAnswer(res []int) int {

	answer := 0

	for _, v := range res {
		answer += v
	}

	answer = answer % 10000

	return answer
}
