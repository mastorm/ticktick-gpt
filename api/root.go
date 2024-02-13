package api

import "net/http"

func ServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", handleAuthorization())

	return mux
}
