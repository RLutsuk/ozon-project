package repository

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
)

type RepositoryI interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}
