package inmemoryrep

import (
	"context"
	"sync"

	"github.com/RLutsuk/ozon-project/graph/model"
)

type inmemoryStore struct {
	comment map[string]*model.Comment
	mu      sync.RWMutex
}

func New() *inmemoryStore {
	return &inmemoryStore{
		comment: make(map[string]*model.Comment),
	}
}

func (r *inmemoryStore) CreateComment(ctx context.Context, comment *model.Comment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.comment[comment.ID] = comment
	return nil
}

func (r *inmemoryStore) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var comments []*model.Comment
	for _, com := range r.comment {
		comments = append(comments, com)
	}
	return comments, nil
}
