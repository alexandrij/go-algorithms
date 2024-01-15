package sort

import (
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	{
		sl := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
		expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		Sort(sl)

		if !slices.Equal(sl, expect) {
			t.Errorf("expect sorted sl to equal %v, but %v", expect, sl)
		}
	}

	{
		sl := []int{20, 1, 19, 2, 18, 3, 17, 4, 16, 5, 15, 6, 14, 7, 13, 8, 12, 9, 11, 10}
		expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

		Sort(sl)

		if !slices.Equal(sl, expect) {
			t.Errorf("expect sorted sl to equal %v, but %v", expect, sl)
		}
	}

	{
		sl := []string{"h", "g", "f", "e", "d", "c", "b", "a"}
		expect := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

		Sort(sl)

		if !slices.Equal(sl, expect) {
			t.Errorf("expect sorted sl to equal %v, but %v", expect, sl)
		}
	}
}
