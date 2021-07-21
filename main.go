package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/DSEMSec/get-subdomains/source"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	host := os.Getenv(source.DBHost)
	port := os.Getenv(source.DBPort)
	user := os.Getenv(source.DBUser)
	pass := os.Getenv(source.DBPass)

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Error().Err(err).Msg("database port number invalid")
		panic("database port number invalid")
	}

	log.Info().Msg("connecting database...")
	repo, err := source.NewRepo(host, portInt, user, pass)
	if err != nil {
		log.Error().Err(err).Msg("cannot initialize repository")
		panic("cannot initialize repository")
	}

	serverPort := os.Getenv(source.ServerPort)
	serverPortInt, err := strconv.Atoi(serverPort)
	if err != nil {
		log.Error().Err(err).Msg("server port number invalid")
		panic("server port number invalid")
	}

	router := source.NewRouter(repo)
	log.Info().Msg("server started")
	http.ListenAndServe(fmt.Sprintf(":%v", serverPortInt), router)
}
