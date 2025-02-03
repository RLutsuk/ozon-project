package graph

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	usecasecomment "github.com/RLutsuk/ozon-project/internal/comment/usecase"
)

type CommentResolver struct {
	commentUC usecasecomment.UseCaseI
}

func (r *CommentResolver) CreateCommentResolver(ctx context.Context, comment *model.Comment) error {
	return r.commentUC.CreateComment(ctx, comment)
}

func (r *CommentResolver) GetCommentResolver(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
	return r.commentUC.GetComments(ctx, postId, limit, offset)
}

func New(commentUC usecasecomment.UseCaseI) *CommentResolver {
	return &CommentResolver{
		commentUC: commentUC,
	}
}
