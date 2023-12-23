package home

import (
	"net/http"

	"github.com/minoritea/chat/container"
)

type Container = *container.Container

func GetHandler(c Container) http.HandlerFunc {
	return c.GetTemplateRenderer().CompileHTTPHandler("home", nil, http.StatusOK)
}
