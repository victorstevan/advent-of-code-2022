package utils

import (
	"reflect"
	"testing"
)

func TestTopKValues(t *testing.T) {
	input := []int{1, 2, 3, 4, 10, 33}

	result := TopKValues(input, 3)

	expectedResult := []int{4, 10, 33}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("TopKValues(input, 3) = %v, want %v", result, expectedResult)
	}
}
