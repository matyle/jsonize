package jsonize

import (
	"fmt"
	"testing"
)

type A struct {
	A string
	B int
}

func TestV(t *testing.T) {
	a := A{
		A: "a",
		B: 1,
	}
	b := A{
		A: "b",
		B: 2,
	}

	datas := []A{a, b}
	dataMap := map[string]A{
		"jack": a,
		"mary": b,
	}
	fmt.Println(V(datas, true))
	fmt.Println(V(dataMap, true))
}
