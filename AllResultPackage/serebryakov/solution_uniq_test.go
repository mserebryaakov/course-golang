package serebryakov_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mserebryaakov/course-golang/build/serebryakov"
)

func TestUniq(t *testing.T) {
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := serebryakov.SolutionUniq(correct)
	if !reflect.DeepEqual(result, 9) {
		t.Fatal("Incorrect result")
	}

	correct2 := []int{}
	result2 := serebryakov.SolutionUniq(correct2)
	if !reflect.DeepEqual(result2, 0) {
		fmt.Print(result2)
		t.Fatal("Incorrect result2")
	}

	correct3 := []int{1, 2, 3, 3, 3, 3, 3, 2, 1}
	result3 := serebryakov.SolutionUniq(correct3)
	if !reflect.DeepEqual(result3, 0) {
		t.Fatal("Incorrect result3")
	}
}
