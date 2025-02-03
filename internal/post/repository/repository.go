package repository

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
)

type RepositoryI interface {
	CreatePost(ctx context.Context, post *model.Post) error
	GetPostByID(ctx context.Context, id string) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
}
