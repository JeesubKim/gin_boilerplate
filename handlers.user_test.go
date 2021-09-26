package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	r := initRoutes()
	req, _ := http.NewRequest("POST", "/user/login", bytes.NewBuffer([]byte(`{"id":"test","password":"test123"}`)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"id":"test","password":"test123"}`, w.Body.String())

}

// func userLogout(t *testing.T){
//
// }
// func userRegister(t *testing.T){
//
// }
