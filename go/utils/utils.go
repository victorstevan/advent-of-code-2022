package utils

import (
	"container/heap"
	"fmt"
	"os"
)

func ReadInputFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", fmt.Errorf("Failed to read from file: %w", err)
	}

	return string(data), nil
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TopKValues(arr []int, k int) []int {
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}

	for i := k; i < len(arr); i++ {
		if arr[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}

	topK := make([]int, k)
	for i := 0; i < k; i++ {
		topK[i] = heap.Pop(h).(int)
	}

	return topK
}
