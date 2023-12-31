package template

import (
	"bytes"
	"database/sql"
	"embed"
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	//go:embed *.tmpl
	tmplFS embed.FS
)

type Container interface {
	GetDB() *sql.DB
}

type Renderer struct {
	tmpl *template.Template
}

func NewRenderer() (*Renderer, error) {
	tmpl := template.New("").Funcs(helpers)
	tmpl, err := tmpl.ParseFS(tmplFS, "*.tmpl")
	if err != nil {
		return nil, err
	}
	return &Renderer{tmpl: tmpl}, nil
}

func (r *Renderer) RenderHTML(w http.ResponseWriter, name string, data any, code int) {
	const suffix = ".html.tmpl"
	var buf bytes.Buffer
	err := r.tmpl.ExecuteTemplate(&buf, name+suffix, data)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_, err = io.Copy(w, &buf)
	if err != nil {
		log.Println(err)
	}
}

func (r *Renderer) RenderOkHTML(w http.ResponseWriter, name string, data any) {
	r.RenderHTML(w, name, data, http.StatusOK)
}

func (r *Renderer) CompileHTTPHandler(name string, data any, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		r.RenderHTML(w, name, data, code)
	}
}

func (r *Renderer) RenderStream(w http.ResponseWriter, name string, data any, code int) {
	const suffix = ".stream.tmpl"
	var buf bytes.Buffer
	err := r.tmpl.ExecuteTemplate(&buf, name+suffix, data)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/vnd.turbo-stream.html; charset=utf-8")
	w.WriteHeader(code)
	_, err = io.Copy(w, &buf)
	if err != nil {
		log.Println(err)
	}
}
