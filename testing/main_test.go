package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップします")
	}

	v := IsOne(i)

	if !v {
		t.Errorf("%v != 1", i)
	}

	// if !IsOne(1) {
	// 	t.Fatal("failed test")
	// }
	// if IsOne(0) {
	// 	t.Fatal("failed test")
	// }
}
