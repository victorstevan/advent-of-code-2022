package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/victorstevan/advent-of-code-2022/utils"
)

func SplitSections(str string) []string {
	str = strings.ReplaceAll(str, "\r\n", "\n") // Normalize newlines
	str = strings.TrimSpace(str)                // Remove leading and trailing whitespace
	sections := strings.Split(str, "\n\n")      // Split on blank lines

	for i, section := range sections {
		sections[i] = strings.ReplaceAll(section, "\n", " ")
	}

	return sections
}

func SumSections(sections []string) ([]int, error) {
	sectionsSum := make([]int, 0)

	for _, section := range sections {
		fields := strings.Fields(section)
		count := 0

		for _, field := range fields {
			value, err := strconv.Atoi(field)

			if err != nil {
				return nil, fmt.Errorf("Failed to parse string to integer: %w", err)
			}

			count += value
		}

		sectionsSum = append(sectionsSum, count)
	}

	return sectionsSum, nil
}

func MaxValueOfArray(arr []int) int {
	maxVal := int(math.Inf(-1))

	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func SumArray(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {
	input, readErr := utils.ReadInputFile("./input.txt")
	if readErr != nil {
		log.Fatalln(readErr)
	}

	sections := SplitSections(input)

	summedValuesForSections, sumErr := SumSections(sections)
	if sumErr != nil {
		log.Fatalln(sumErr)
	}

	maxValue := MaxValueOfArray(summedValuesForSections)
	topValues := utils.TopKValues(summedValuesForSections, 3)
	topValuesSummed := SumArray(topValues)

	fmt.Println("Part 1 - Result: ", maxValue)
	fmt.Println("Part 2 - Result: ", topValuesSummed)
}
