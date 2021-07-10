package main

import (
	"log"
	"logic_king_backend/cores"
)

func main() {
	cfg := cores.NewConfig()
	mongodb := cores.NewMongo(&cores.MongoOption{
		Username: cfg.Username(),
		Password: cfg.Password(),
		DBName:   cfg.DBName(),
		URL:      cfg.DBHost(),
		Port:     cfg.DBPort(),
	})

	migration := NewMigration()

	migration.AddBatch(NewCreateProjectCollectionMigrationBatch(mongodb))

	err := migration.Up()
	if err != nil {
		log.Fatal(err)
	}
}
