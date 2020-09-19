package main

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const URL = "https://httpmock.test.com/users"

func getUser() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
func TestHttpMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", URL,
		httpmock.NewStringResponder(http.StatusOK, `{"id": 1, "name": "Bond"}`))

	err := getUser()
	assert.NoError(t, err)

	info := httpmock.GetCallCountInfo()

	key := "GET" + " " + URL

	assert.Equal(t, 1, info[key])
}
