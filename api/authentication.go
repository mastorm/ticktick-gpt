package api

import "net/http"

func handleAuthorization() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/oauth/callback", handleOauthCallback)

	return mux
}

func handleOauthCallback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
