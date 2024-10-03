package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashboard_ItRedirectsGuestsToLoginPage(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	res, err := tc.Client.Get(tc.Server.URL + "/dashboard")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusFound, res.Request.Response.StatusCode)
	assert.Equal(t, "/login", res.Request.Response.Header.Get("Location"))
}
