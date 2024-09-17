package main

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	go main()
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code for / is wrong. Have: %d, want: %d.", resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
}
