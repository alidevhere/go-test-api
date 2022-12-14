package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	r := SetupRouter("test")

	//Test 1
	u := url.Values{}
	u.Add("name", "Ali")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.URL.RawQuery = u.Encode()
	r.ServeHTTP(w, req)
	// Result 1
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `"Hello Ali !!"`, w.Body.String())

	//Test 2
	req2, _ := http.NewRequest("GET", "/login", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	//Result 2
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, `"Hello Annonymous user !!"`, w2.Body.String())
}
