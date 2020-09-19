package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	err       error
	userEmpty User
	readed    Readed
	user      = User{
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
)

func TestAssert(t *testing.T) {
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 0, userEmpty.ID)
	assert.NotEqual(t, 100500, userEmpty.ID)
	assert.Empty(t, userEmpty)

	assert.ElementsMatch(t, user.Books, []string{
		"The Place in Dalhousie",
		"Catch and Kill",
		"Kit’s Wilderness",
		"Peak",
		"The Secret Garden",
	})

	go func() {
		for i := 0; i < 100; i++ {
			readed.M.Lock()
			readed.Number++
			readed.M.Unlock()

			time.Sleep(1 * time.Millisecond)
		}
	}()

	assert.Eventually(t, func() bool {
		readed.M.Lock()
		defer readed.M.Unlock()

		if readed.Number > 2 {
			return true
		}

		return false
	}, 50*time.Millisecond, 1*time.Millisecond)
}
