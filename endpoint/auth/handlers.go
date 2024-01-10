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
		data.Flashes = session.MustGetFlashes(c, w, r)
		data.AssetPath = c.Config().AssetPath()
		renderer.RenderOkHTML(w, "auth", data)
	}
}

func PostHandler(c Container) http.HandlerFunc {
	oauth2Config := c.Config().GithubOAuth2Config()
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
	oauth2Config := c.Config().GithubOAuth2Config()
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		sessionState, ok := session.MustGet(c, r).Values["oauth2state"].(string)
		if !ok || sessionState != state {
			log.Println("invalid oauth2 state")
			session.RedirectWithErrorFlash(c, w, r, "/", "Invalid oauth2 state")
			return
		}

		code := r.URL.Query().Get("code")
		ctx := context.WithValue(r.Context(), oauth2.HTTPClient, client)
		token, err := oauth2Config.Exchange(ctx, code)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(c, w, r, "/", "Failed to exchange oauth2 code")
			return
		}

		githubUser, err := getGithubUser(ctx, token)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(c, w, r, "/", "Failed to get github user")
			return
		}
		sessionUser, err := user.FindOrCreateUser(r.Context(), c, githubUser.Login)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(c, w, r, "/", "Internal server error")
			return
		}
		err = session.StoreNewSession(r.Context(), c, w, r, sessionUser.ID)
		if err != nil {
			log.Println(err)
			session.RedirectWithErrorFlash(c, w, r, "/", "Internal server error")
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

type githubUser struct {
	Login string `json:"login"`
}

var client = &http.Client{
	Transport: transport,
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
