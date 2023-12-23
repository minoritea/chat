package signin

import (
	"net/http"

	"github.com/minoritea/chat/container"
)

type Container = *container.Container

func GetHandler(c Container) http.HandlerFunc {
	return c.GetTemplateRenderer().CompileHTTPHandler("signin", nil, http.StatusOK)
}

func PostHandler(c Container) http.HandlerFunc {
	renderer := c.GetTemplateRenderer()
	return func(w http.ResponseWriter, r *http.Request) {
		account := r.PostFormValue("account")
		if account == "" {
			var data struct{ Error string }
			data.Error = "Account name is required"
			renderer.RenderHTML(w, "signin", data, http.StatusBadRequest)
			return
		}

		password := r.PostFormValue("password")
		if password == "" {
			var data struct{ Error string }
			data.Error = "Password is required"
			renderer.RenderHTML(w, "signin", data, http.StatusBadRequest)
			return
		}

		var data struct{ Error string }
		data.Error = "Failed to sign in"
		renderer.RenderHTML(w, "signin", data, http.StatusBadRequest)
	}
}
