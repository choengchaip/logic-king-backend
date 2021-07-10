package projects

import (
	"go.mongodb.org/mongo-driver/bson"
	"logic_king_backend/cores"
	"logic_king_backend/models"
)

type IProjectService interface {
	Pagination() ([]*models.Project, error)
}

type ProjectService struct {
	ctx *cores.AppContext
}

func NewProjectService(ctx *cores.AppContext) *ProjectService {
	return &ProjectService{
		ctx: ctx,
	}
}

func (s *ProjectService) Pagination() ([]*models.Project, error) {
	items := make([]*models.Project, 0)
	cur, err := s.ctx.Mongo.Database.Collection(ProjectCollectionName).Find(s.ctx.Mongo.GetContext(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.All(s.ctx.Mongo.GetContext(), &items)
	if err != nil {
		return nil, err
	}

	return items, err
}
