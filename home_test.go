package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	res, err := tc.Client.Get(tc.Server.URL + "/")

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)
}
