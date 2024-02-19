package api

import (
	"github.com/mastorm/ticktick-gpt/ticktick"
	"log"
	"net/http"
)

type AuthRoute struct {
	app *ticktick.Application
}

func (a *AuthRoute) handleAuthorization(tickTickApp *ticktick.Application) *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/callback", a.handleOauthCallback)
	mux.HandleFunc("/auth/redirect", a.handleInitiateAuthFlow)
	return mux
}

func (a *AuthRoute) handleInitiateAuthFlow(w http.ResponseWriter, r *http.Request) {
	url, err := a.app.GetAuthorizeUrl("lul", "http://localhost:8080/oauth/callback")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	log.Println(*url)

	http.Redirect(w, r, *url, 303)
}

func (a *AuthRoute) handleOauthCallback(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, World!"))
}
