package main

import (
	"fmt"
	"sort"
)

// SortingStrategy is a strategy interface for sorting
type SortingStrategy interface {
	Sort([]int) []int
}

// BubbleSort implements SortingStrategy using Bubble sort algorithm
type BubbleSort struct{}

func (b *BubbleSort) Sort(data []int) []int {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

// QuickSort implements SortingStrategy using Quick sort algorithm
type QuickSort struct{}

func (q *QuickSort) Sort(data []int) []int {
	sort.Ints(data)
	return data
}

// Sorter has a SortingStrategy
type Sorter struct {
	strategy SortingStrategy
}

func (s *Sorter) Sort(data []int) []int {
	return s.strategy.Sort(data)
}

func main() {
	data := []int{5, 2, 4, 1, 3}
	sorter := &Sorter{new(BubbleSort)}
	sortedData := sorter.Sort(data)
	fmt.Println("BubbleSort result:", sortedData)

	data = []int{5, 2, 4, 1, 3}
	sorter = &Sorter{new(QuickSort)}
	sortedData = sorter.Sort(data)
	fmt.Println("QuickSort result:", sortedData)
}
