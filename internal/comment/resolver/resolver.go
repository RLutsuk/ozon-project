package graph

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	usecasecomment "github.com/RLutsuk/ozon-project/internal/comment/usecase"
)

type CommentResolver struct {
	commentUC usecasecomment.UseCaseI
}

func (r *CommentResolver) CreateCommentResolver(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error) {
	return r.commentUC.CreateComment(ctx, input)
}

func (r *CommentResolver) GetUserByID(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return r.commentUC.GetUserByID(ctx, obj)
}

// func (r *CommentResolver) GetCommentResolver(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
// 	return r.commentUC.GetComments(ctx, postId, limit, offset)
// }

func New(commentUC usecasecomment.UseCaseI) *CommentResolver {
	return &CommentResolver{
		commentUC: commentUC,
	}
}
