package arrayandslices

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	t.Run("Slice of 10 numbers", func(t *testing.T) {
		nums := []int{1, 2, 4, 56, 32, 12, 44, 78, 1, 2}
		got := SumSlice(nums)
		want := 232
		if got != want {
			t.Errorf("got %d, want %d given, %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3, 4}, []int{3, 5, 6}, []int{1, 4, 3})
	want := []int{8, 14, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestTailAll(t *testing.T) {
	got := TailAll([]int{2, 3, 1}, []int{3, 4, 5}, []int{})
	want := []int{1, 5, -1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}

}
