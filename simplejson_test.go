package main

import (
	"testing"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleJson(t *testing.T) {
	var userJson = `{
			"id":       123456,
			"name":     "Rauf",
			"age":      30,
			"password": "qwerty123",
			"email":    "user@test.ru",
			"books": [
				"Peak"
			],
			"weight": 77.5
		}`

	js, err := simplejson.NewJson([]byte(userJson))
	assert.NoError(t, err)

	assert.Equal(t, js.Get("id").MustInt64(), int64(123456))
	assert.Equal(t, js.Get("name").MustString(), "Rauf")

	arr, err := js.Get("books").Get("array").Array()
	assert.NotEqual(t, nil, arr)
	for _, v := range arr {
		r, ok := v.(string)
		require.Equal(t, true, ok)
		assert.Equal(t, "Peak", r)
	}
}
