package inmemoryrep

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/RLutsuk/ozon-project/graph/model"
	"github.com/google/uuid"
)

type inmemoryStore struct {
	posts map[string]*model.Post
	mu    sync.RWMutex
}

func New() *inmemoryStore {
	return &inmemoryStore{
		posts: make(map[string]*model.Post),
	}
}

func (r *inmemoryStore) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	post.Created = time.Now()
	post.ID = uuid.New().String()
	r.posts[post.ID] = post
	return post, nil
}

func (r *inmemoryStore) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	post, exists := r.posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (r *inmemoryStore) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var posts []*model.Post
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}
