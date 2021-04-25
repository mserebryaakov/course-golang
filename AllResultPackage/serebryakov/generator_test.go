package serebryakov_test

import (
	"reflect"
	"testing"

	"github.com/mserebryaakov/course-golang/build/serebryakov"
)

func TestGenerator(t *testing.T) {
	result := serebryakov.SolutionSquareGenerator(0, 5)
	if !reflect.DeepEqual(result, []int{1, 4, 9, 16}) {
		t.Fatal("Incorrect result")
	}

	result2 := serebryakov.SolutionSquareGenerator(0, 3)
	if !reflect.DeepEqual(result2, []int{1, 4}) {
		t.Fatal("Incorrect result2")
	}

	result3 := serebryakov.SolutionSquareGenerator(2, 0)
	if !reflect.DeepEqual(result3, []int{}) {
		t.Fatal("Incorrect result3")
	}

	result4 := serebryakov.SolutionSquareGenerator(-2, 4)
	if !reflect.DeepEqual(result4, []int{1}) {
		t.Fatal("Incorrect result4")
	}
}
