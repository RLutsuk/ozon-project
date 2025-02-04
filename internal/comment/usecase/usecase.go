package usecase

import (
	"context"
	"fmt"

	"github.com/RLutsuk/ozon-project/graph/model"
	commentrep "github.com/RLutsuk/ozon-project/internal/comment/repository"
	"github.com/sirupsen/logrus"
)

type UseCaseI interface {
	CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error)
	// GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error)
}

type useCase struct {
	commentRepository commentrep.RepositoryI
}

func New(commentRepository commentrep.RepositoryI) UseCaseI {
	return &useCase{
		commentRepository: commentRepository,
	}
}

func (uc *useCase) CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error) {
	var newcomment model.Comment
	newcomment.Body = input.Body
	newcomment.PostId = input.PostID
	newcomment.UserId = input.UserID

	if input.ParentID != "" {
		newcomment.ParentId = input.ParentID
		parentComment, _ := uc.commentRepository.GetCommentByiD(ctx, input.ParentID)
		newcomment.Level = parentComment.Level + 1
		fmt.Println(newcomment.Level)
		if newcomment.Level == 1 {
			newcomment.RootID = input.ParentID
		} else {
			newcomment.RootID = parentComment.RootID
		}
	}

	comment, err := uc.commentRepository.CreateComment(ctx, &newcomment)
	if err == nil {
		logrus.Info("Comment succesfully created")
	}
	return comment, err
}

// func (uc *useCase) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
// 	rootcomments, _, err := uc.commentRepository.GetComments(ctx, postId, limit, offset)
// 	if err == nil {
// 		logrus.Info("Comment succesfully found")
// 	}
// 	return rootcomments, err
// }
