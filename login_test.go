package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin_Page(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	res, err := tc.Client.Get(tc.Server.URL + "/login")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
