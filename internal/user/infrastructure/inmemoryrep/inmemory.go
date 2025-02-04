package inmemoryrep

import (
	"context"
	"errors"
	"sync"

	"github.com/RLutsuk/ozon-project/graph/model"
)

type inmemoryStore struct {
	users map[string]*model.User
	mu    sync.RWMutex
}

func New() *inmemoryStore {
	return &inmemoryStore{
		users: make(map[string]*model.User),
	}
}

func (r *inmemoryStore) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
