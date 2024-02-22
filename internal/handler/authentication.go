package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) TickTickAuth(c echo.Context) (err error) {
	url, err := h.TodoApp.GetAuthorizeUrl("lul")
	if err != nil {
		return err
	}

	return c.Redirect(303, *url)

}

func (h *Handler) TickTickAuthCallback(c echo.Context) error {
	code := c.QueryParam("code")
	res, err := h.TodoApp.RequestAccessToken(code)
	if err != nil {
		return err
	}
	return c.String(200, res.AccessToken)
}
