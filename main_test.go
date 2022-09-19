package main

import (
	"testing"

	"github.com/rjeczalik/notify"
)

func TestConstant(t *testing.T) {
	if notify.Create == notify.Remove {
		t.FailNow()
	}
}
