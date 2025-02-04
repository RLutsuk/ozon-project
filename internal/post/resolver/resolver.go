package graph

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	usecasepost "github.com/RLutsuk/ozon-project/internal/post/usecase"
)

type PostResolver struct {
	postUC usecasepost.UseCaseI
}

func (r *PostResolver) CreatePostResolver(ctx context.Context, inputPost model.CreatePostInput) (*model.Post, error) {
	return r.postUC.CreatePost(ctx, inputPost)
}

func (r *PostResolver) GetAllPostsResolver(ctx context.Context) ([]*model.Post, error) {
	return r.postUC.GetAllPosts(ctx)
}

func (r *PostResolver) GetPostResolver(ctx context.Context, id string, limit, offset *int32) (*model.Post, error) {
	return r.postUC.GetPost(ctx, id, limit, offset)
}

func (r *PostResolver) GetUserByID(ctx context.Context, obj *model.Post) (*model.User, error) {
	return r.postUC.GetUserByID(ctx, obj)
}

func New(postUC usecasepost.UseCaseI) *PostResolver {
	return &PostResolver{
		postUC: postUC,
	}
}
