package lowercase

import (
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Lowercase{})
}

type Lowercase struct{}

func (Lowercase) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.lowercase",
		New: func() caddy.Module { return new(Lowercase) },
	}
}

func (h Lowercase) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	lowerPath := strings.ToLower(r.URL.Path)
	if r.URL.Path != lowerPath {
		// Optionally redirect, or rewrite internally
		r.URL.Path = lowerPath
	}
	return next.ServeHTTP(w, r)
}
