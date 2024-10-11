package main

import (
	"container/heap"
	"fmt"
)

// tipagem de um heap de inteiros
type IntegerHeap []int

// IntegerHeap method - get the length of Heap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// check if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

// swaps the element of i to j index (algorithm pattern)
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

// method heap - pushes the item onto the heap
func (iheap *IntegerHeap) Push(x interface{}) {
	*iheap = append(*iheap, x.(int))
}

// method heap - pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	previous := *iheap
	n := len(previous)
	x1 := previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	// initializing the heap with values
	intHeap := &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
