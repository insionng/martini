package martini

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GzipAll(t *testing.T) {
	// Set up
	recorder := httptest.NewRecorder()
	before := false

	m := New()
	m.Use(Gzip())
	m.Use(func(r http.ResponseWriter) {
		r.(ResponseWriter).Before(func(rw ResponseWriter) {
			before = true
		})
	})

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(recorder, r)

	// Make our assertions
	_, ok := recorder.HeaderMap[HeaderContentEncoding]
	if ok {
		t.Error(HeaderContentEncoding + " present")
	}

	ce := recorder.Header().Get(HeaderContentEncoding)
	if strings.EqualFold(ce, "gzip") {
		t.Error(HeaderContentEncoding + " is 'gzip'")
	}

	recorder = httptest.NewRecorder()
	r.Header.Set(HeaderAcceptEncoding, "gzip")
	m.ServeHTTP(recorder, r)

	// Make our assertions
	_, ok = recorder.HeaderMap[HeaderContentEncoding]
	if !ok {
		t.Error(HeaderContentEncoding + " not present")
	}

	ce = recorder.Header().Get(HeaderContentEncoding)
	if !strings.EqualFold(ce, "gzip") {
		t.Error(HeaderContentEncoding + " is not 'gzip'")
	}

	if before == false {
		t.Error("Before hook was not called")
	}
}
