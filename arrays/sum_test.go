package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 2})
		want := []int{2, 11}

		checkSums(t, got, want)
	})

	t.Run("safely runs on empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1})
		want := []int{0, 0}

		checkSums(t, got, want)
	})
}
