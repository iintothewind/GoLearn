package linq

import (
	"fmt"
	"testing"

	"github.com/ahmetalpbalkan/go-linq"
)

func TestSelect01(t *testing.T) {
	squares := []int{}

	linq.
		Range(1, 10).
		Select(func(x interface{}) interface{} { return x.(int) * x.(int) }).
		ToSlice(&squares)

	fmt.Println(squares)
}
