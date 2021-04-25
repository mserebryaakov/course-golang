package serebryakov_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mserebryaakov/course-golang/build/serebryakov"
)

func TestRotate(t *testing.T) {
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := serebryakov.SolutionRotate(correct, 3)
	if !reflect.DeepEqual(result, []int{7, 8, 9, 1, 2, 3, 4, 5, 6}) {
		t.Fatal("Incorrect result")
	}

	correct2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result2 := serebryakov.SolutionRotate(correct2, 1)
	if !reflect.DeepEqual(result2, []int{9, 1, 2, 3, 4, 5, 6, 7, 8}) {
		fmt.Print(result2)
		t.Fatal("Incorrect result2")
	}

	correct3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result3 := serebryakov.SolutionRotate(correct3, 0)
	if !reflect.DeepEqual(result3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal("Incorrect result3")
	}
}
