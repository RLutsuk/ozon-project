package inmemoryrep

import (
	"context"
	"sync"

	"github.com/RLutsuk/ozon-project/graph/model"
	"github.com/pkg/errors"
)

type inmemoryStore struct {
	users map[string]*model.User
	mu    sync.RWMutex
}

// Сразу добавляются тестовые данные
func New() *inmemoryStore {
	return &inmemoryStore{
		users: map[string]*model.User{
			"11111111-1111-1111-1111-111111111111": {
				ID:        "11111111-1111-1111-1111-111111111111",
				Username:  "alice",
				Email:     "alice@example.com",
				Firstname: "Alice",
				Lastname:  "Johnson",
			},
			"22222222-2222-2222-2222-222222222222": {
				ID:        "22222222-2222-2222-2222-222222222222",
				Username:  "bob",
				Email:     "bob@example.com",
				Firstname: "Bob",
				Lastname:  "Smith",
			},
			"33333333-3333-3333-3333-333333333333": {
				ID:        "33333333-3333-3333-3333-333333333333",
				Username:  "charlie",
				Email:     "charlie@example.com",
				Firstname: "Charlie",
				Lastname:  "Brown",
			},
		},
	}
}

func (r *inmemoryStore) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if user, exists := r.users[id]; !exists {
		return nil, errors.Wrap(model.ErrUserNotFound, "database error: user not found (method GetUserByID, table users)")
	} else {
		return user, nil
	}
}

func (r *inmemoryStore) GetUsers(ctx context.Context, id []string) ([]*model.User, error) {
	var users []*model.User
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, idUser := range id {
		users = append(users, r.users[idUser])
	}
	return users, nil
}
