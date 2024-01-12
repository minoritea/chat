package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
	"github.com/minoritea/chat/resource"
	"github.com/oklog/ulid/v2"
	"golang.org/x/oauth2"
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
	oauth2Config := c.Config().GithubOAuth2Config()
	return func(w http.ResponseWriter, r *http.Request) {
		state := ulid.Make().String()
		s := session.MustGet(c, r)
		s.Values["oauth2state"] = state
		session.MustSave(s, r, w)
		http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusSeeOther)
	}
}

func GetCallbackHandler(c Container) http.HandlerFunc {
	oauth2Config := c.Config().GithubOAuth2Config()
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		s := session.MustGet(c, r)
		sessionState, ok := s.Values["oauth2state"].(string)
		if !ok || sessionState != state {
			log.Println("invalid oauth2 state")
			session.RedirectWithErrorFlash(w, r, s, "/", "Invalid oauth2 state")
			return
		}
		delete(s.Values, "oauth2state")

		code := r.URL.Query().Get("code")
		token, err := oauth2Config.Exchange(r.Context(), code)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Failed to exchange oauth2 code")
			return
		}

		githubUser, err := getGithubUser(r.Context(), token)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(w, r, s, "/", "Failed to get github user")
			return
		}
		sessionUser, err := user.FindOrCreateUser(r.Context(), c, githubUser.Login)
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

type githubUser struct {
	Login string `json:"login"`
}

func getGithubUser(ctx context.Context, token *oauth2.Token) (*githubUser, error) {
	const endpoint = "https://api.github.com/user"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	token.SetAuthHeader(req)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user githubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
