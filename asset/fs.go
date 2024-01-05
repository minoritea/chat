package asset

import "embed"

//go:embed js/*.js css/*.css
var FS embed.FS
