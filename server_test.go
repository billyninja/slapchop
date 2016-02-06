package main

import "testing"

func TestSum(t *testing.T) {
	if 1+1 != 2 {
		t.Error(" 1 + 1 != 2")
	}
}
