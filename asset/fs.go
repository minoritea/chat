package asset

import "embed"

//go:embed js/*.js css/*.css favicon.ico
var FS embed.FS
