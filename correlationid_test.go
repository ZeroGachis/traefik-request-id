package correlationid

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddCorrelationIdInHeaderIfNoneExist(t *testing.T) {
	cfg := CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "correlation-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	cid := req.Header.Get("X-Correlation-ID")
	if cid == "" {
		t.Errorf("FAK")
	}
}

func TestKeepCorrelationIdInHeaderIfOneExist(t *testing.T) {
	cfg := CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "correlation-id-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	existingCid := "some-existing-correlation-id"
	req.Header.Set("X-Correlation-ID", existingCid)

	handler.ServeHTTP(recorder, req)

	cid := req.Header.Get("X-Correlation-ID")
	if cid != existingCid {
		t.Errorf("FAK")
	}
}
