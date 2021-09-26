package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandingPageRoute(t *testing.T) {
	router := initRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"What is":"gin.H?"}`, w.Body.String())
}

func TestLoginPageRoute(t *testing.T) {
	router := initRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"login":"api"}`, w.Body.String())
}

func TestSignupPageRoute(t *testing.T) {
	router := initRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/signup", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"signup":"api"}`, w.Body.String())
}
