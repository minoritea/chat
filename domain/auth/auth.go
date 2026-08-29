// Package auth provides domain logic for OAuth2 authentication.
package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/minoritea/chat/config"
	"github.com/oklog/ulid/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var oAuth2StateKey = "oauth2state"

// Token wraps oauth2.Token.
type Token interface {
	SetAuthHeader(req *http.Request)
}

// OAuth2 wraps oauth2.Config with the state and token used in the authentication flow.
type OAuth2 struct {
	oauth2.Config
	state *string
	token Token
}

// ConfigContainer provides a config.Config.
type ConfigContainer interface {
	Config() config.Config
}

// New returns a new OAuth2.
func New(c ConfigContainer) *OAuth2 {
	oauth2Config := oauth2.Config{
		ClientID:     c.Config().GithubClientID,
		ClientSecret: c.Config().GithubClientSecret,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
	return &OAuth2{Config: oauth2Config}
}

// SetNewState sets a new state to the session and o.
func (o *OAuth2) SetNewState(s *sessions.Session) *OAuth2 {
	state := ulid.Make().String()
	s.Values[oAuth2StateKey] = state
	o.state = &state
	return o
}

// WithStateFromSession sets the state from the session to o.
func (o *OAuth2) WithStateFromSession(s *sessions.Session) *OAuth2 {
	state, ok := s.Values[oAuth2StateKey].(string)
	if ok {
		o.state = &state
	}
	return o
}

// ValidateState reports whether state matches the stored state.
func (o *OAuth2) ValidateState(state string) bool {
	return o.state != nil && *o.state == state
}

// DeleteState deletes the state from the session.
func (o *OAuth2) DeleteState(s *sessions.Session) {
	delete(s.Values, oAuth2StateKey)
}

// AuthCodeURL returns the URL to redirect the user for GitHub authentication.
func (o *OAuth2) AuthCodeURL() string {
	if o.state == nil {
		panic("state is not set")
	}

	return o.Config.AuthCodeURL(*o.state)
}

// Exchange exchanges the auth code for a token and stores it in o.
func (o *OAuth2) Exchange(ctx context.Context, code string) error {
	token, err := o.Config.Exchange(ctx, code)
	if err != nil {
		return err
	}
	o.token = token
	return nil
}

// GetGithubUserName fetches the GitHub account name of the authenticated user.
func (o *OAuth2) GetGithubUserName(ctx context.Context) (string, error) {
	if o.token == nil {
		return "", errors.New("token is not set")
	}

	const endpoint = "https://api.github.com/user"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	o.token.SetAuthHeader(req)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	var user struct {
		Login string `json:"login"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", err
	}
	return user.Login, nil
}
