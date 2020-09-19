package main

import (
	"testing"

	"github.com/go-test/deep"
)

var (
	userEqual = User{
		ID:       123456,
		Name:     "Rauf",
		Age:      30,
		Password: "qwerty123",
		Email:    "user@test.ru",
		Books: []string{
			"The Place in Dalhousie",
			"Catch and Kill",
			"Kit’s Wilderness",
			"Peak",
			"The Secret Garden",
		},
		Weight: 77.5,
	}

	userNotEqual = User{
		ID:       456,
		Name:     "Bond",
		Age:      30,
		Password: "123456",
		Email:    "user@test.ru",
		Books: []string{
			"The Place in Dalhousie",
		},
		Weight: 77.5,
	}
)

func TestDeep(t *testing.T) {
	if diff := deep.Equal(user, userEqual); diff != nil {
		t.Error(diff)
	}

	if diff := deep.Equal(user, userNotEqual); diff != nil {
		t.Error(diff)
	}
}
