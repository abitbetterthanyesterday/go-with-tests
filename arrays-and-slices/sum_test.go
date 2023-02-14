package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func (t *testing.T){

	numbers := []int{1,2,3,4,5}

	got := Sum(numbers)
	expected := 15

	if got != expected { t.Errorf("given '%v' got '%d' expected '%d'", numbers, got, expected)
	}
	})
}

func TestSumAll(t *testing.T){

	got :=SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	checkSums(got, want, t)
}

func checkSums(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' expected '%v'", got, want)
	}
}

func TestSumAllTails(t *testing.T){
	t.Run("make the sums of some slices", func (t *testing.T){

	got := SumAllTails([]int{1,2}, []int{0,9})
	want := []int{2,9}
	checkSums(got, want, t)
	})

	t.Run("safely sum empty slices", func (t *testing.T){
	got := SumAllTails([]int{}, []int{0,9})
	want := []int{0,9}
	checkSums(got, want, t)

	})
}