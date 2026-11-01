package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func get(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	newHandler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestHealth(t *testing.T) {
	rec := get(t, "/health")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"ok"`) {
		t.Fatalf("body = %s", rec.Body)
	}
}

func TestHello(t *testing.T) {
	cases := []struct{ path, want string }{
		{"/api/hello", "hello, world"},
		{"/api/hello?name=vunet", "hello, vunet"},
		{"/api/hello?name=" + strings.Repeat("a", 500), "hello, " + strings.Repeat("a", 64)},
	}
	for _, c := range cases {
		var body map[string]string
		if err := json.Unmarshal(get(t, c.path).Body.Bytes(), &body); err != nil {
			t.Fatal(err)
		}
		if body["message"] != c.want {
			t.Errorf("GET %.40s: got %q, want %q", c.path, body["message"], c.want)
		}
	}
}

func TestSecurityHeaders(t *testing.T) {
	rec := get(t, "/health")
	for k, v := range securityHeaders {
		if got := rec.Header().Get(k); got != v {
			t.Errorf("%s = %q, want %q", k, got, v)
		}
	}
}

func TestUnknownRouteAndMethod(t *testing.T) {
	if rec := get(t, "/nope"); rec.Code != http.StatusNotFound {
		t.Errorf("unknown route: %d", rec.Code)
	}
	rec := httptest.NewRecorder()
	newHandler().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/health", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST /health: %d", rec.Code)
	}
}
