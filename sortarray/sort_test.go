package sortarray

import (
	"reflect"
	"testing"
)

type TestTable struct {
	input  []int
	output []int
}

func TestSortArray(t *testing.T) {
	tests := []TestTable{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{6, 5, 3, 1, 3, 1, 5}, []int{1, 1, 3, 3, 5, 5, 6}},
	}

	for _, test := range tests {
		test.output = SortArray(test.input)
		if !equal(test.input, test.output) {
			t.Errorf("got %v want %v", test.input, test.output)
		}
	}
}

func equal(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}
