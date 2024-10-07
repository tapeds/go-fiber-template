package migrations

import (
	"log"

	"github.com/tapeds/go-fiber-template/entity"
	"gorm.io/gorm"
)

func dropAllTables(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		&entity.User{},
	); err != nil {
		return err
	}

	log.Println("All tables dropped successfully.")
	return nil
}

func Fresh(db *gorm.DB) error {
	if err := dropAllTables(db); err != nil {
		log.Printf("Error dropping tables: %v", err)
		return err
	}

	if err := Migrate(db); err != nil {
		log.Printf("Error during migration: %v", err)
		return err
	}

	return nil
}
