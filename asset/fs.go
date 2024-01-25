package asset

import "embed"

//go:embed js/*.js js/*.js.map css/*.css favicon.ico
var FS embed.FS
