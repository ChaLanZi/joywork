package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestProxyHeaders(t *testing.T) {
	assert := assert2.New(t)
	testHeaders := map[string]string{
		"foo": "bar",
		"a":   "test",
		"1":   "2 3 4",
	}
	req, err := http.NewRequest("GET", "https://www.joywork.com/foo/bar/test.jpg?local=true", nil)
	assert.NoError(err)
	for k, v := range testHeaders {
		req.Header.Set(k, v)
	}
	rec := httptest.NewRecorder()
	ProxyHeaders(req.Header, rec.Header())
	for k, v := range testHeaders {
		assert.Equal(v, rec.Header().Get(k))
	}
	assert.Equal(len(testHeaders), len(rec.Header()))
}
