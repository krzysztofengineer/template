package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	Server *httptest.Server
	Client *http.Client
	DB     *sql.DB
}

func NewTestCase() *TestCase {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	s := httptest.NewServer(NewApp(db))

	return &TestCase{
		Server: s,
		Client: s.Client(),
		DB:     db,
	}
}

func (tc *TestCase) Close() {
	tc.Server.Close()
}
