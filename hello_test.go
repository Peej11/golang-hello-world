package main

import (
    "testing"
    "net/http/httptest"
    "io/ioutil"
    "net/http"
)

func TestHello(t *testing.T) {
    expected := "Hello, James!"
    req := httptest.NewRequest(http.MethodGet, "/James", nil)
    w := httptest.NewRecorder()
    HelloServer(w, req)
    res := w.Result()
    defer res.Body.Close()

    data, err := ioutil.ReadAll(res.Body)
    
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    
    if string(data) != expected {
        t.Errorf("Expected '" + expected + "' but got '%v'", string(data))
    }
}