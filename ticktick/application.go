package ticktick

import (
	"net/url"
	"strings"
)

const (
	ScopesTasksRead  = "tasks:read"
	ScopesTasksWrite = "tasks:write"
)

const (
	Url = "https://ticktick.com"
)

type Application struct {
	ClientId string
	Scopes   []string
}

const (
	clientId     = "client_id"
	Scope        = "scope"
	RedirectUri  = "redirect_uri"
	State        = "state"
	ResponseType = "response_type"
)

// GetAuthorizeUrl returns the URL to redirect the user to in order to authorize the application
func (app *Application) GetAuthorizeUrl(state string, redirectUri string) (*string, error) {
	u, err := url.Parse(Url)
	if err != nil {
		return nil, err
	}

	u.Path = "/oauth/authorize"

	params := url.Values{}
	params.Add("scope", strings.Join(app.Scopes, " "))
	params.Add("client_id", app.ClientId)
	params.Add("state", state)
	params.Add("redirect_uri", redirectUri)
	params.Add("response_type", "code")

	u.RawQuery = params.Encode()
	result := u.String()

	return &result, nil
}
