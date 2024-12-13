package sort

import "testing"

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1, 0, 0}, []int{0, 0, 1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		result := quickSort(test.input)
		if len(result) != len(test.expect) {
			t.Errorf("quickSort(%v) => %v, expect %v", test.input, result, test.expect)
			continue
		}

		for i := range result {
			if result[i] != test.expect[i] {
				t.Errorf("quickSort(%v) => %v, expect %v", test.input, result, test.expect)
				break
			}
		}
	}
}

func quickSort(slice []int) (result []int) {
	if len(slice) <= 1 {
		return slice
	}

	mid := len(slice) / 2
	var left []int
	var right []int
	for _, s := range slice {
		if s < slice[mid] {
			left = append(left, s)
		} else if s > slice[mid] {
			right = append(right, s)
		}
	}

	result = append(result, quickSort(left)...)
	result = append(result, slice[mid])
	result = append(result, quickSort(right)...)

	return
}
