//go:build !js

package auth

import "net/http"

var transport http.RoundTripper = http.DefaultTransport
