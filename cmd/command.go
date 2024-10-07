package cmd

import (
	"log"
	"os"

	"github.com/tapeds/go-fiber-template/migrations"
	"gorm.io/gorm"
)

func Commands(db *gorm.DB) {
	migrate := false
	seed := false
	fresh := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
		if arg == "--migrate-fresh" {
			fresh = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if seed {
		if err := migrations.Seeder(db); err != nil {
			log.Fatalf("error migration seeder: %v", err)
		}
		log.Println("seeder completed successfully")
	}
	if fresh {
		if err := migrations.Fresh(db); err != nil {
			log.Fatalf("error migration fresh: %v", err)
		}
		log.Println("fresh migration completed successfully")
	}
}
