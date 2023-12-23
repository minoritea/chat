package signup

import (
	"net/http"

	"github.com/minoritea/chat/container"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
)

type Container = container.Container

func GetHandler(c *Container) http.HandlerFunc {
	return c.GetTemplateRenderer().CompileHTTPHandler("signup", nil, http.StatusOK)
}

func PostHandler(c *Container) http.HandlerFunc {
	renderer := c.GetTemplateRenderer()
	return func(w http.ResponseWriter, r *http.Request) {
		account := r.PostFormValue("account")
		if account == "" {
			var data struct{ Error string }
			data.Error = "Account name is required"
			renderer.RenderHTML(w, "signup", data, http.StatusBadRequest)
			return
		}

		password := r.PostFormValue("password")
		if password == "" {
			var data struct{ Error string }
			data.Error = "Password is required"
			renderer.RenderHTML(w, "signup", data, http.StatusBadRequest)
			return
		}

		sessionUser, err := user.RegisterUser(c, account, password)
		if err != nil {
			var data struct{ Error string }
			data.Error = "Sign up failed"
			renderer.RenderHTML(w, "signup", data, http.StatusInternalServerError)
			return
		}

		err = session.StoreNewSession(c, w, r, sessionUser.ID)
		if err != nil {
			var data struct{ Error string }
			data.Error = "Sign up failed"
			renderer.RenderHTML(w, "signup", data, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
