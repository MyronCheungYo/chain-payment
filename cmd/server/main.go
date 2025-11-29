package main

import (
	"log"

	"backend/pkg/server"
)

func main() {
	cfg := server.LoadConfigFromEnv()
	srv, err := server.New(cfg)
	if err != nil {
		log.Fatalf("init server: %v", err)
	}

	log.Printf("listening on %s", cfg.ListenAddr)
	if err := srv.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
