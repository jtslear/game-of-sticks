package main

import (
	"testing"
)

func TestValidateStickCount(t *testing.T) {
	if !validateStickCount(2, 1, 3) {
		t.Error("Expected to be true")
	}
	if validateStickCount(1, 2, 3) {
		t.Error("Expected to be true")
	}
	if validateStickCount(3, 1, 2) {
		t.Error("Expected to be true")
	}
}

func TestSwapPlayer(t *testing.T) {
	control := []string{"one", "two"}
	test := swapPlayer(control)

	if test[0] != "two" {
		t.Error("Swapped incorrectly")
	}
	if test[1] != "one" {
		t.Error("Swapped incorrectly")
	}
}

