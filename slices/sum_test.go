package slices

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("Got %d, expected %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Test the sum of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}

	t.Run("Test summing tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("Test summing tails of slices with more than two numbers", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 4})
		want := []int{5, 13}
		checkSums(t, got, want)
	})

	t.Run("Test summing tail of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
