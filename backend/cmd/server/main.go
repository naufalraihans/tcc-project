package main

import (
	"context"
	"log"

	"tcc-itpln/backend/config"
	"tcc-itpln/backend/internal/router"
	"tcc-itpln/backend/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	if err := router.New(cfg, db).Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
