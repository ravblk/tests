package main

import (
	"testing"

	"go.uber.org/goleak"
)

func Myfunc() {
	go func() {
	}()
}

func MyfuncWithLeak() {
	go func() {
		for {
		}
	}()
}

func TestLeak(t *testing.T) {
	t.Run("not leak", func(t *testing.T) {
		defer goleak.VerifyNone(t)
		Myfunc()
	})

	t.Run("leak", func(t *testing.T) {
		defer goleak.VerifyNone(t)
		MyfuncWithLeak()
	})
}
