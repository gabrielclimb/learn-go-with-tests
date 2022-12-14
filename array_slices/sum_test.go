package array_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbersSizeDefined := []int{1, 2, 3, 4, 5}

		got := Sum(numbersSizeDefined)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbersSizeDefined)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbersSizeNotDefined := []int{1, 2, 3}

		got := Sum(numbersSizeNotDefined)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbersSizeNotDefined)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("return total in each slice passed", func(t *testing.T) {
		array1 := []int{1, 2, 3}
		array2 := []int{0, 4, 5}

		got := SumAll(array1, array2)
		want := []int{6, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
