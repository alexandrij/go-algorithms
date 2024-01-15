package sort

import (
	"math"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

const maxInsertion = 12

func insertionSort[T Ordered](sl []T) {
	for i, l := 1, len(sl); i < l; i++ {
		for j := i; j > 0 && sl[j-1] > sl[j]; j-- {
			swap(sl, j, j-1)
		}
	}
}

func swap[T Ordered](sl []T, i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

func quickSort[T Ordered](sl []T, start, end int) {
	if len(sl) < 2 {
		return
	}

	left := start
	right := end
	pivot := sl[int(math.Floor(float64((start+end)/2)))]

	for left <= right {
		for sl[left] < pivot {
			left++
		}
		for sl[right] > pivot {
			right--
		}

		if left <= right {
			swap(sl, left, right)
			left++
			right--
		}
	}

	if left < end {
		quickSort(sl, left, end)
	}

	if right > start {
		quickSort(sl, start, right)
	}
}

func Sort[T Ordered](sl []T) {
	if len(sl) < maxInsertion {
		insertionSort(sl)
	} else {
		quickSort(sl, 0, len(sl)-1)
	}
}
