package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/minoritea/chat/resource"
	"github.com/minoritea/chat/domain/session"
	"github.com/minoritea/chat/domain/user"
	"github.com/oklog/ulid/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Container = resource.Container

var port = "8080"
var oauth2Config = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Endpoint:     github.Endpoint,
	// RedirectURL:  "http://localhost:" + port + "/oauth/callback",
	Scopes: []string{"user:email"},
}

func GetHandler(c Container) http.HandlerFunc {
	renderer := c.Renderer()
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct{ Flashes []session.Flash }
		data.Flashes = session.MustGetFlashes(c, w, r)
		renderer.RenderOkHTML(w, "auth", data)
	}
}

func PostHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := ulid.Make().String()
		s := session.MustGet(c, r)
		s.Values["oauth2state"] = state
		err := s.Save(r, w)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusSeeOther)
	}
}

func GetCallbackHandler(c Container) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		sessionState, ok := session.MustGet(c, r).Values["oauth2state"].(string)
		if !ok || sessionState != state {
			log.Println("invalid oauth2 state")
			session.MustAddFlash(c, w, r, session.NewErrorFlash("invalid oauth2 state"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		code := r.URL.Query().Get("code")
		token, err := oauth2Config.Exchange(r.Context(), code)
		if err != nil {
			log.Println(err)
			session.MustAddFlash(c, w, r, session.NewErrorFlash("failed to exchange oauth2 code"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		githubUser, err := getGithubUser(r.Context(), token)
		if err != nil {
			log.Println(err)
			session.MustAddFlash(c, w, r, session.NewErrorFlash("failed to get github user"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		sessionUser, err := user.FindOrCreateUser(r.Context(), c, githubUser.Login)
		if err != nil {
			log.Println(err)
			session.MustAddFlash(c, w, r, session.NewErrorFlash("internal server error"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		err = session.StoreNewSession(r.Context(), c, w, r, sessionUser.ID)
		if err != nil {
			log.Println(err)
			session.MustAddFlash(c, w, r, session.NewErrorFlash("internal server error"))
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
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
