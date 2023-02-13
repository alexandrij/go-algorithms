package algorithms

import (
	"fmt"
	"sort"
)

func InstersectsTwoSlice(sl1 []int, sl2 []int) []int {
	i := 0
	j := 0
	sl := []int{}

	sort.Ints(sl1)
	sort.Ints(sl2)

	for i < len(sl1) && j < len(sl2) {
		for sl1[i] < sl2[j] && i < len(sl1) {
			i++
		}
		for sl2[j] < sl1[i] && j < len(sl2) {
			j++
		}
		if sl1[i] == sl2[j] {
			sl = append(sl, sl1[i])
		}
		i++
		j++
	}
	return sl
}

func InstersectsTwoSliceCached(sl1 []int, sl2 []int) []int {
	var sl []int
	cached := make(map[int]int)

	for i := range sl1 {
		cached[sl1[i]]++
	}

	for i := range sl2 {
		if count, ok := cached[sl2[i]]; ok {
			if count > 0 {
				sl = append(sl, sl2[i])
				cached[sl2[i]]--
			}
		}
	}
	return sl
}

func main() {
	fmt.Println(InstersectsTwoSlice([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(InstersectsTwoSlice([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))

	fmt.Println(InstersectsTwoSliceCached([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(InstersectsTwoSliceCached([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
