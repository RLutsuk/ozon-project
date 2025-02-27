package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"errors"

	"github.com/RLutsuk/ozon-project/graph/model"
	userloaders "github.com/RLutsuk/ozon-project/internal/user/dataloader"
	"github.com/sirupsen/logrus"
)

// User is the resolver for the user field.
func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	// user, err := r.CommentResolver.GetUserByID(ctx, obj)
	// if err != nil {
	// 	logrus.WithError(err).Error("comment resolver error: get user")
	// 	switch {
	// 	case errors.Is(err, model.ErrUserNotFound):
	// 		return nil, model.ErrUserNotFound
	// 	default:
	// 		return nil, model.ErrInternalServer
	// 	}
	// }
	// return user, nil

	//dataloader
	user, err := userloaders.GetUser(ctx, obj.UserId)
	if err != nil {
		logrus.WithError(err).Error("comment resolver error: get user")
		switch {
		case errors.Is(err, model.ErrUserNotFound):
			return nil, model.ErrUserNotFound
		default:
			return nil, model.ErrInternalServer
		}
	}
	return user, err
}

// Level is the resolver for the level field.
func (r *commentResolver) Level(ctx context.Context, obj *model.Comment) (int32, error) {
	level := int32(obj.Level)
	return level, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*model.Post, error) {
	post, err := r.PostResolver.CreatePostResolver(ctx, input)
	if err != nil {
		logrus.WithError(err).Error("mutation resolver error: create post")
		switch {
		case errors.Is(err, model.ErrBadData):
			return nil, model.ErrBadData
		case errors.Is(err, model.ErrUserNotFound):
			return nil, model.ErrUserNotFound
		default:
			return nil, model.ErrInternalServer
		}
	}
	return post, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error) {
	comment, err := r.CommentResolver.CreateCommentResolver(ctx, input)
	if err != nil {
		logrus.WithError(err).Error("mutation resolver error: create comment")
		switch {
		case errors.Is(err, model.ErrBadData):
			return nil, model.ErrBadData
		case errors.Is(err, model.ErrCommentNotFound):
			return nil, model.ErrCommentNotFound
		case errors.Is(err, model.ErrUserNotFound):
			return nil, model.ErrUserNotFound
		case errors.Is(err, model.ErrPostNotFound):
			return nil, model.ErrPostNotFound
		case errors.Is(err, model.ErrСommentsProhibited):
			return nil, model.ErrСommentsProhibited
		default:
			return nil, model.ErrInternalServer
		}
	}
	return comment, nil
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	// user, err := r.PostResolver.GetUserByID(ctx, obj)
	// if err != nil {
	// 	logrus.WithError(err).Error("post resolver error: get user")
	// 	switch {
	// 	case errors.Is(err, model.ErrUserNotFound):
	// 		return nil, model.ErrUserNotFound
	// 	default:
	// 		return nil, model.ErrInternalServer
	// 	}
	// }
	// return user, err

	//dataloader
	user, err := userloaders.GetUser(ctx, obj.UserID)
	if err != nil {
		logrus.WithError(err).Error("post resolver error: get user")
		switch {
		case errors.Is(err, model.ErrUserNotFound):
			return nil, model.ErrUserNotFound
		default:
			return nil, model.ErrInternalServer
		}
	}
	return user, err
}

// Getpost is the resolver for the getpost field.
func (r *queryResolver) Getpost(ctx context.Context, id string, limit *int32, offset *int32) (*model.Post, error) {
	post, err := r.PostResolver.GetPostResolver(ctx, id, limit, offset)
	if err != nil {
		logrus.WithError(err).Error("query resolver error: get post")
		switch {
		case errors.Is(err, model.ErrPostNotFound):
			return nil, model.ErrPostNotFound
		default:
			return nil, model.ErrInternalServer
		}
	}
	return post, nil
}

// Getposts is the resolver for the getposts field.
func (r *queryResolver) Getposts(ctx context.Context) ([]*model.Post, error) {
	posts, err := r.PostResolver.GetAllPostsResolver(ctx)
	if err != nil {
		logrus.WithError(err).Error("query resolver error: get all posts")
		switch {
		case errors.Is(err, model.ErrPostsDontExist):
			return nil, model.ErrPostsDontExist
		default:
			return nil, model.ErrInternalServer
		}
	}
	return posts, nil
}

// NewCommentToPost is the resolver for the newCommentToPost field.
func (r *subscriptionResolver) NewCommentToPost(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	// commentEvent := make(chan *model.Comment, 1)
	// go func() {
	// 	<-ctx.Done()
	// }()
	// return commentEvent, nil
	return nil, nil
}

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
