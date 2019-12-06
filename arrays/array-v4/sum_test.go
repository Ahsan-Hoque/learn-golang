package main

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d but want %d for numbers %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d but want %d", got, want)
	}
}
