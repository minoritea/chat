// Package asset provides embedded static assets.
package asset

import "embed"

// FS holds the embedded static assets.
//
//go:embed js/*.js js/*.js.map css/*.css favicon.ico
var FS embed.FS
