package main

import (
	"fmt"
	"testing"
)

func TestPermutateIntCount(t *testing.T) {
	got := PermutateNum(3).Count()
	want := 6

	if got != want {
		t.Errorf("Expect %d got %d", want, got)
	}

	got2 := PermutateNum(10).Count()
	want2 := 3628800

	if got2 != want2 {
		t.Errorf("Expect %d got %d", want2, got2)
	}
}

func TestPermutateSliceCount(t *testing.T) {
	data := []interface{}{1, 2, 3}
	got := PermutateSlice(data).Count()
	want := 6

	if got != want {
		t.Errorf("Expect %d got %d", want, got)
	}

}

func TestPermutateList(t *testing.T) {
	r := PermutateNum(5).List()
	fmt.Println(r)
}
