package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRoomHandle(t *testing.T) {
	router := setUpRouter()
	path := "/room"
	method := "POST"

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("TestCreateRoomHandle, room was not created: expected %v, have %v!\n",
			http.StatusCreated, rr.Code)
	}
}
