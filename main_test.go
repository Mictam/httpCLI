package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestIndexHandlerOkStatus(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}

func TestIndexHandler404Status(t *testing.T) {
	HtmlPath = "!@#404@#!"
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Unexpected status code %d", resp.StatusCode)
	}
}