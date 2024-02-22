package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mastorm/ticktick-gpt/internal/handler"
	"github.com/mastorm/ticktick-gpt/ticktick"
	"log"
	"os"
)

const (
	EnvClientId     = "CLIENT_ID"
	EnvClientSecret = "CLIENT_SECRET"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := ticktick.Application{
		ClientId:     os.Getenv(EnvClientId),
		ClientSecret: os.Getenv(EnvClientSecret),
		Scopes:       []string{ticktick.ScopesTasksRead, ticktick.ScopesTasksWrite},
		RedirectUri:  "http://localhost:8080/oauth/callback",
	}

	h := handler.Handler{
		TodoApp: app,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/auth/redirect", h.TickTickAuth)
	e.GET("/oauth/callback", h.TickTickAuthCallback)

	e.Logger.Fatal(e.Start(":8080"))
}
