package inmemoryrep

import (
	"context"
	"sync"
	"time"

	"github.com/RLutsuk/ozon-project/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type inmemoryStore struct {
	posts map[string]*model.Post
	mu    sync.RWMutex
}

func New() *inmemoryStore {
	return &inmemoryStore{
		posts: map[string]*model.Post{
			"aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1": {
				ID:            "aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1",
				Title:         "Hello World",
				Body:          "This is my first post on this platform!",
				UserID:        "11111111-1111-1111-1111-111111111111",
				Allowcomments: true,
				Created:       time.Now(),
			},
			"aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2": {
				ID:            "aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2",
				Title:         "Test Post",
				Body:          "This is my second post!",
				UserID:        "11111111-1111-1111-1111-111111111111",
				Allowcomments: false,
				Created:       time.Now(),
			},
			"aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3": {
				ID:            "aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3",
				Title:         "Tech Trends 2025",
				Body:          "A deep dive into upcoming technologies.",
				UserID:        "22222222-2222-2222-2222-222222222222",
				Allowcomments: true,
				Created:       time.Now(),
			},
			"aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4": {
				ID:            "aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4",
				Title:         "No Comments Allowed",
				Body:          "This post is for reading only.",
				UserID:        "33333333-3333-3333-3333-333333333333",
				Allowcomments: false,
				Created:       time.Now(),
			},
		},
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
		return nil, errors.Wrap(model.ErrPostNotFound, "database error: post not found (method GetPostByID, table posts)")
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
	if len(posts) == 0 {
		return nil, errors.Wrap(model.ErrPostsDontExist, "database error: posts don't exist (method GetAllPosts, table posts)")
	}
	return posts, nil
}
