package main

import "testing"

func TestCreateANewDB(t *testing.T) {

	if 1 != 0 {
		t.Errorf("Could not create a new DB instance")
	}
}
