package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPing
//  GET /ping
func TestPing(t *testing.T) {
	ts := httptest.NewServer(Server())
	defer ts.Close()
	url := ts.URL + "/ping"
	res, err := http.Get(url)
	Must(err)
	message, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	Must(err)
	log.Printf("%s", message)
}
