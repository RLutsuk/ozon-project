package inmemoryrep

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/RLutsuk/ozon-project/graph/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type inmemoryStore struct {
	// commentsByPost map[string][]*model.Comment
	commentsByID map[string]*model.Comment
	mu           sync.RWMutex
}

func New() *inmemoryStore {
	return &inmemoryStore{
		// commentsByPost: make(map[string][]*model.Comment),
		commentsByID: make(map[string]*model.Comment),
	}
}

func (r *inmemoryStore) CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	comment.Created = time.Now()
	comment.ID = uuid.New().String()
	if comment.ParentId == "" {
		comment.Level = 0
	} else {
		parentCom := r.commentsByID[comment.ParentId]
		comment.Level = parentCom.Level + 1
		if comment.Level == 1 {
			comment.RootID = comment.ParentId
		} else {
			comment.RootID = parentCom.RootID
		}
	}

	/*
		if comment.ParentId == "" {
			comment.Level = 0
			r.commentsByPost[comment.PostId] = append(r.commentsByPost[comment.PostId], comment)
		} else {
			parent, exists := r.commentsByID[comment.ParentId]
			if !exists {
				return nil, fmt.Errorf("parent comment not found")
			}
			comment.Level = parent.Level + 1
			if parent.Replies == nil {
				parent.Replies = make([]*model.Comment, 0)
			}
			parent.Replies = append(parent.Replies, comment)
		}
		r.commentsByID[comment.ID] = comment*/
	r.commentsByID[comment.ID] = comment
	return comment, nil
}

func (r *inmemoryStore) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, []*model.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var parentcomments []*model.Comment
	var childcomments []*model.Comment
	for _, comment := range r.commentsByID {
		if comment.PostId == postId {
			if comment.Level == 0 {
				parentcomments = append(parentcomments, comment)
			} else {
				childcomments = append(childcomments, comment)
			}
		}
	}
	sort.Slice(parentcomments, func(i, j int) bool {
		return parentcomments[i].Created.Before(parentcomments[j].Created)
	})
	// if len(parentcomments) == 0 {
	// 	return nil, nil, errors.Wrap(model.ErrCommentsDontExist, "database error (method GetComments, table comments)")
	// }
	start := offset
	end := offset + limit
	if start > len(parentcomments) {
		return nil, nil, errors.Wrap(model.ErrCommentOffset, "database error: comments don't exist (method GetComments, table comments)")
	}
	if end > len(parentcomments) {
		end = len(parentcomments)
	}

	return parentcomments[start:end], childcomments, nil
}

func (r *inmemoryStore) GetCommentByID(ctx context.Context, commentID string) (*model.Comment, error) {
	if comment, ok := r.commentsByID[commentID]; !ok {
		return nil, errors.Wrap(model.ErrCommentNotFound, "database error: comment not found (method GetCommentByiD, table comments)")
	} else {
		return comment, nil
	}
}
