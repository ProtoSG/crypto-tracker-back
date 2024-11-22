package services

import (
	"fmt"
	"log"
	"os"
)

type Env struct {
	URL string
}

func NewEnv() *Env {

	databaseURL, ok := getEnvVar("TURSO_DATABASE_URL")
	if !ok {
		log.Fatal("TURSO_DATABASE_URL not found in .env file")
	}

	authToken, ok := getEnvVar("TURSO_AUTH_TOKEN")
	if !ok {
		log.Fatal("TURSO_AUTH_TOKEN not found in .env file")
	}

	url := fmt.Sprintf("%s?authToken=%s", databaseURL, authToken)

	return &Env{URL: url}
}

func getEnvVar(key string) (string, bool) {
	value := os.Getenv(key)
	if value == "" {
		return "", false
	}
	return value, true
}
