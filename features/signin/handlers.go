package signin

import (
	"net/http"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
)

type Container = container.Container

func GetHandler(c *Container) http.HandlerFunc {
	return c.GetTemplateRenderer().CompileHTTPHandler("signin", nil, http.StatusOK)
}

func PostHandler(c *Container) http.HandlerFunc {
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

		sessionUser, err := user.GetByAccoutNameAndPassword(r.Context(), c, account, password)
		if err == nil {
			var data struct{ Error string }
			data.Error = "Sign in failed"
			renderer.RenderHTML(w, "signin", data, http.StatusBadRequest)
			return
		}
		session.StoreNewSession(r.Context(), c, w, r, sessionUser.ID)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
