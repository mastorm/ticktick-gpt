package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mastorm/ticktick-gpt/api"
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
		Scopes:   []string{ticktick.SCOPES_TASK_READ, ticktick.SCOPES_TASK_WRITE},
	}
	fmt.Println(app.ClientId)

	mux := api.ApiServeMux()

	http.ListenAndServe(":8080", mux)
}
