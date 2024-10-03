package main

import (
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	Server *httptest.Server
	Client *http.Client
}

func NewTestCase() *TestCase {
	s := httptest.NewServer(NewApp())

	return &TestCase{
		Server: s,
		Client: s.Client(),
	}
}

func (tc *TestCase) Close() {
	tc.Server.Close()
}
