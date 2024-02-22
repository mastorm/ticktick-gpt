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
	ClientId     string
	ClientSecret string
	Scopes       []string
	RedirectUri  string
}

func getUrl(path string) (*url.URL, error) {
	u, err := url.Parse(Url)
	if err != nil {
		return nil, err
	}
	u.Path = path
	return u, nil
}

func formatScopes(app *Application) string {
	return strings.Join(app.Scopes, " ")
}
