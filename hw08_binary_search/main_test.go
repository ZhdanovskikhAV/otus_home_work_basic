package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		expect int
	}{
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 11, 5},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 3, 1},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 1, 0},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 19, 9},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 20, -1},
		{[]int{}, -1, -1},
	}

	for _, test := range tests {
		result := binarySearch(test.arr, test.target)
		assert.Equal(t, test.expect, result)
	}
}
