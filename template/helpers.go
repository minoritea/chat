// Package template provides HTML template rendering and helper functions.
package template

import (
	"html/template"
	"time"
)

var helpers = template.FuncMap{"formatAsDateTime": FormatAsDateTime}

// FormatAsDateTime formats t using the time.DateTime layout.
func FormatAsDateTime(t time.Time) string {
	return t.In(time.Local).Format(time.DateTime)
}
