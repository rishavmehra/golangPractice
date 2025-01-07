package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{23, 1, 10, 5, 2}
	got := MergeSort(arr)
	want := []int{1, 2, 5, 10, 23}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got is %v, want is %v", got, want)
	}

}