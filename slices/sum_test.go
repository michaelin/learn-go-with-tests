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
		got := SumAll( []int{1, 2}, []int{0, 9} )
		want := []int{3, 9}

		if !reflect.DeepEqual(want,got) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}