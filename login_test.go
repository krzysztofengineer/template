package main

import (
	"io"
	"net/http"
	"net/url"
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

func TestLogin_FormValidation(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	res, err := tc.Client.PostForm(tc.Server.URL+"/login", nil)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	assert.Contains(t, string(body), "The email is required")
}

func TestLogin_InvalidEmailFormat(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	f := url.Values{}
	f.Set("email", "invalid")

	res, err := tc.Client.PostForm(tc.Server.URL+"/login", f)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	assert.Contains(t, string(body), "The email format is invalid")
}

func TestLogin_UserIsCreated(t *testing.T) {
	tc := NewTestCase()
	defer tc.Close()

	f := url.Values{}
	f.Set("email", "valid@example.com")

	tc.Client.PostForm(tc.Server.URL+"/login", f)

	var exists bool
	err := tc.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = 'valid@example.com')", "valid@example.com").Scan(&exists)
	assert.Nil(t, err)
	assert.True(t, exists)
}
