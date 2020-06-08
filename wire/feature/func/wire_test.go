package build

import "testing"

func TestFunc(t *testing.T) {
	initEndingA, fc, err := InitEndingA("mario", "kitty")
	if err != nil {
		t.Error(err)
		return
	}
	initEndingA.Appear()
	fc()
}
