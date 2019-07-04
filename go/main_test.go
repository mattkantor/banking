package main

import "testing"

func TestAnything(t *testing.T) {

	if 1 != 0 {
		t.Errorf("1 does not equal zero, fix your code")
	}
}
