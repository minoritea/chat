package auth

import (
	"log"
	"net/http"

	"github.com/minoritea/chat/domain/auth"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
	"github.com/minoritea/chat/resource"
)

type Container = resource.Container

func GetHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			session.FlashData
			AssetPath string
		}
		s, err := session.Get(c, r)
		if err != nil {
			s = session.MustNew(c, r)
		}
		data.Flashes = session.GetFlashes(s)
		data.AssetPath = c.Config().AssetPath()
		session.MustSave(s, r, w)
		renderer.RenderOkHTML(w, "auth", data)
	}
}

func PostHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := session.MustGet(c, r)
		o := auth.New(c).SetNewState(s)
		session.MustSave(s, r, w)
		http.Redirect(w, r, o.AuthCodeURL(), http.StatusSeeOther)
	}
}

func GetCallbackHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		s := session.MustGet(c, r)
		o := auth.New(c).WithStateFromSession(s)
		if !o.ValidateState(state) {
			log.Println("invalid oauth2 state")
			session.RedirectWithErrorFlash(w, r, s, "/", "Invalid oauth2 state")
			return
		}
		o.DeleteState(s)

		code := r.URL.Query().Get("code")
		err := o.Exchange(r.Context(), code)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Failed to exchange oauth2 code")
			return
		}

		userName, err := o.GetGithubUserName(r.Context())
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Failed to get github user")
			return
		}
		sessionUser, err := user.FindOrCreateUser(r.Context(), c, userName)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Internal server error")
			return
		}
		dbSession, err := session.PerpetuateSession(r.Context(), c, sessionUser.ID)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Internal server error")
			return
		}
		session.SetSessionID(s, dbSession.ID)
		session.MustSave(s, r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
