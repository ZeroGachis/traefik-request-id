package traefik_request_id

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddRequestIdInHeaderIfNoneExist(t *testing.T) {
	cfg := CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "sw-request-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	cid := req.Header.Get("X-Request-ID")
	if cid == "" {
		t.Errorf("Request ID has not been generated")
	}
}

func TestKeepRequestIdInHeaderIfOneExist(t *testing.T) {
	cfg := CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "sw-request-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	existingCid := "some-existing-request-id"
	req.Header.Set("X-Request-ID", existingCid)

	handler.ServeHTTP(recorder, req)

	cid := req.Header.Get("X-Request-ID")
	if cid != existingCid {
		t.Errorf("Existing Request ID has not been kept")
	}
}
