package ticktick

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// GetAuthorizeUrl returns the URL to redirect the user to in order to authorize the application
func (app *Application) GetAuthorizeUrl(state string) (*string, error) {
	u, err := getUrl("/oauth/authorize")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("scope", formatScopes(app))
	params.Add("client_id", app.ClientId)
	params.Add("state", state)
	params.Add("redirect_uri", app.RedirectUri)
	params.Add("response_type", "code")

	u.RawQuery = params.Encode()
	result := u.String()

	return &result, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (app *Application) RequestAccessToken(code string) (*TokenResponse, error) {
	u, err := getUrl("/oauth/token")
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Add("code", code)
	form.Add("grant_type", "authorization_code")
	form.Add("scopes", formatScopes(app))
	form.Add("redirect_uri", app.RedirectUri)

	c := http.Client{}

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(app.ClientId, app.ClientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	var body TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
