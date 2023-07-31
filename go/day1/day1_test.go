package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/victorstevan/advent-of-code-2022/utils"
)

func TestSplitSections(t *testing.T) {
	input, _ := utils.ReadInputFile("./test_input.txt")

	result := SplitSections(input)
	expectedResult := []string{"1000 2000 3000", "4000", "5000 6000", "7000 8000 9000", "10000"}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("SplitSections(input) = %v; want %v", result, expectedResult)
	}

}

func TestSumSections(t *testing.T) {
	input := []string{"1000 2000 3000", "4000", "5000 6000", "7000 8000 9000", "10000"}

	summedSections, _ := SumSections(input)

	expectedResult := []int{6000, 4000, 11000, 24000, 10000}

	if !reflect.DeepEqual(summedSections, expectedResult) {
		t.Errorf("SumSections(input) = %v, want %v", summedSections, expectedResult)
	}

}

func TestMaxValueOfArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 10}

	result := MaxValueOfArray(input)

	if result != 10 {
		t.Errorf("MaxValueOfArray(input) = %d, want 10", result)
	}
}

func TestSumArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 10, 33}

	result := SumArray(input)

	if result != 53 {
		t.Errorf("SumArray(input) = %d, want 53", result)
	}
}

func generateRandomArray(K int) []int {
	// Set a seed for the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create an array to store the random numbers
	randomArray := make([]int, K)

	// Fill the array with random numbers
	for i := 0; i < K; i++ {
		randomArray[i] = rand.Intn(100) // You can change 100 to any range you desire for the random numbers.
	}

	return randomArray
}
