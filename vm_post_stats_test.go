package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	// Top - top command
	Top string = "top"
)

// TestVMUpdateStatus
//  PUT /simagix/:hostname/commands
func TestVMPostStats(t *testing.T) {
	log.Println("TestVMPostStats")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in TestVMPostStats", r)
			t.Fail()
		}
	}()
	ts := httptest.NewServer(Server())
	defer ts.Close()
	hostname := "simagix-analytic"
	var data []byte
	// test POST top
	url := ts.URL + "/simagix/" + hostname + "/commands/" + Top
	data = GetDataFromFile("testdata/" + Top + ".txt")
	postMessage(url, data)
}

func postMessage(url string, data []byte) {
	log.Println(url)
	response, err := http.Post(url, "plain/text", bytes.NewReader(data))
	Must(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	Must(err)
	if response.StatusCode != 201 {
		log.Panic("POST failed", response.StatusCode)
	}
	log.Printf("%s\n", contents)
}
