package main

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestGoFuzz(t *testing.T) {
	f := fuzz.New()

	var (
		object  User
		testInt int
	)

	f.Fuzz(&object)
	f.Fuzz(&testInt)

	assert.NotEmpty(t, object)
	assert.NotEqual(t, 0, testInt)
}
