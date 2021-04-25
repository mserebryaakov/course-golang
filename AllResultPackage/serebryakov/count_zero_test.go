package serebryakov_test

import (
	"reflect"
	"testing"

	"github.com/mserebryaakov/course-golang/build/serebryakov"
)

func TestCountZero(t *testing.T) {
	result := serebryakov.SolutionBinaryGap(0)
	if !reflect.DeepEqual(result, 0) {
		t.Fatal("Incorrect result")
	}

	result2 := serebryakov.SolutionBinaryGap(5)
	if !reflect.DeepEqual(result2, 1) {
		t.Fatal("Incorrect result2")
	}

	result3 := serebryakov.SolutionBinaryGap(7)
	if !reflect.DeepEqual(result3, 0) {
		t.Fatal("Incorrect result3")
	}

	result4 := serebryakov.SolutionBinaryGap(1024)
	if !reflect.DeepEqual(result4, 10) {
		t.Fatal("Incorrect result4")
	}
}
