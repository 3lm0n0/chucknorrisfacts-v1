package main

import (
	"log"

	"github.com/joho/godotenv"
)

func getEnv() (envMap map[string]string, err error) {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return envFile, nil
}

func main() {
	// godotenv package
	envFile, err := getEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables file")
		return
	}
	url := envFile["url"]
	token := envFile["token"]
	// new service
	svc := NewFactService(url, token)
	// new logger
	svc = NewLogger(svc)
	// new server
	server := NewServer(svc)
	// log
	log.Fatal(server.Start(":3000"))
}
