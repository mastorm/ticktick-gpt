package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mastorm/ticktick-gpt/internal/handler"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mastorm/ticktick-gpt/ticktick"
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
		ClientId: os.Getenv(EnvClientId),
		Scopes:   []string{ticktick.ScopesTasksRead, ticktick.ScopesTasksWrite},
	}

	h := handler.Handler{
		TodoApp: app,
	}

	e := echo.New()
	e.GET("/auth/redirect", h.TickTickAuth)
	e.GET("/oauth/callback", h.TickTickAuthCallback)
	
	e.Logger.Fatal(e.Start(":8080"))
}
