package _defer

import (
	"fmt"
	"reflect"
	"testing"
)

type Test struct{ name string }

func (t *Test) Point() {
	fmt.Println(t.name)
}
func Test_defer1(t *testing.T) {
	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}
	for _, t := range ts {
		fmt.Println(reflect.TypeOf(t))
		defer t.Point()
	}
}
