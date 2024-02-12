package api

import "net/http"

func ApiServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", handleAuthorization())

	return mux
}
