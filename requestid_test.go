package requestid_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	plugin "github.com/ZeroGachis/traefik-request-id"
)

func TestAddRequestIdInHeaderIfNoneExist(t *testing.T) {
	cfg := plugin.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := plugin.New(ctx, next, cfg, "sw-request-id-plugin")
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
	cfg := plugin.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := plugin.New(ctx, next, cfg, "sw-request-id-plugin")
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
