package api

import (
	"github.com/mastorm/ticktick-gpt/ticktick"
	"net/http"
)

func ServeMux(app *ticktick.Application) *http.ServeMux {
	mux := http.NewServeMux()

	route := AuthRoute{
		app: app,
	}

	mux.Handle("/", route.handleAuthorization(app))

	return mux
}
