package main

import (
	"testing"
)

func TestPart2(t *testing.T) {
	t.Run("Not nice", func(t *testing.T) {
		if Check2("aaa") {
			t.Error("Test failed, func returned as nice string")
		}
	})

	t.Run("Not nice", func(t *testing.T) {
		if Check2("uurcxstgmygtbstg") {
			t.Error("Test failed, func returned as nice string")
		}
	})

	t.Run("Not nice", func(t *testing.T) {
		if Check2("aaagjfkkeka") {
			t.Error("Test failed, func returned as nice string")
		}
	})

	t.Run("Nice", func(t *testing.T) {
		if !Check2("xxyxx") {
			t.Error("Test failed, func returned as not nice string")
		}
	})

}
