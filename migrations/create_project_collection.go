package main

import (
	"logic_king_backend/cores"
	"logic_king_backend/services/projects"
)

type CreateProjectCollection struct {
	mongo *cores.Mongo
}

func (c *CreateProjectCollection) Name() string {
	return "create_project_collection"
}

func (c *CreateProjectCollection) Up() error {
	return c.mongo.Database.CreateCollection(c.mongo.GetContext(), projects.ProjectCollectionName)
}

func (c *CreateProjectCollection) Down() error {
	return nil
}

func NewCreateProjectCollectionMigrationBatch(mongo *cores.Mongo) IMigrationBatch {
	return &CreateProjectCollection{
		mongo: mongo,
	}
}
