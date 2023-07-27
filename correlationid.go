// Package correlationid a Traefik plugin to add correlation ID to incoming HTTP requests.
package correlationid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Plugin a correlation id plugin.
type Plugin struct {
	next http.Handler
	name string
}

// New created a new Plugin plugin.
func New(_ context.Context, next http.Handler, _ *Config, name string) (http.Handler, error) {
	return &Plugin{
		next: next,
		name: name,
	}, nil
}

func (a *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	cid := req.Header.Get("X-Correlation-ID")
	if cid == "" {
		req.Header.Set("X-Correlation-ID", uuid.New().String())
	}

	a.next.ServeHTTP(rw, req)
}
