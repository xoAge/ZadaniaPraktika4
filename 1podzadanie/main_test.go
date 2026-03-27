package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{5, 3, 8, 4, 2}
	expected := []int{2, 3, 4, 5, 8}

	bubbleSort(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Expected %v, got %v", expected, input)
	}
}
