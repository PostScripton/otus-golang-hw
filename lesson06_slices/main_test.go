package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Напишите функцию `Concat`, которая получает несколько слайсов
// и склеивает их в один длинный.
// { {1, 2, 3}, {4, 5}, {6, 7} }  => {1, 2, 3, 4, 5, 6, 7}

func Concat(slices [][]int) (res []int) {
	for _, slice := range slices {
		res = append(res, slice...)
	}

	return res
}

func TestConcat(t *testing.T) {
	test := []struct {
		slices   [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}

	for _, tc := range test {
		assert.Equal(t, tc.expected, Concat(tc.slices))
	}
}
