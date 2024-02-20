package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) TickTickAuth(c echo.Context) (err error) {
	url, err := h.TodoApp.GetAuthorizeUrl("lul", "http://localhost:8080/oauth/callback")
	if err != nil {
		return err
	}

	return c.Redirect(303, *url)

}

func (h *Handler) TickTickAuthCallback(c echo.Context) error {
	return c.String(200, "Hello, world!")
}
