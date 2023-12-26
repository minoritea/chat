package template

import (
	"html/template"
	"time"
)

var helpers = template.FuncMap{"formatAsDateTime": FormatAsDateTime}

func FormatAsDateTime(t time.Time) string {
	return t.Format(time.DateTime)
}
