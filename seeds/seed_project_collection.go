package main

import (
	"logic_king_backend/cores"
	"logic_king_backend/models"
	"logic_king_backend/services/projects"
)

type SeedProjectCollection struct {
	mongo *cores.Mongo
}

func (c *SeedProjectCollection) Name() string {
	return "seed_project_collection"
}

func (c *SeedProjectCollection) Up() error {
	_, err := c.mongo.Database.Collection(projects.ProjectCollectionName).InsertOne(c.mongo.GetContext(), models.Project{
		Title:              "E-Learngin",
		Description:        "eieijaaa",
		BackgroundImageURL: "not",
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *SeedProjectCollection) Down() error {
	return nil
}

func NewSeedProjectCollectionMigrationBatch(mongo *cores.Mongo) ISeedBatch {
	return &SeedProjectCollection{
		mongo: mongo,
	}
}
