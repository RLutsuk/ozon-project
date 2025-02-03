package repository

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
)

type RepositoryI interface {
	CreateComment(ctx context.Context, comment *model.Comment) error
	GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error)
}
