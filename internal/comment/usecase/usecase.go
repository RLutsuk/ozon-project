package usecase

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	commentrep "github.com/RLutsuk/ozon-project/internal/comment/repository"
)

type UseCaseI interface {
	CreateComment(ctx context.Context, comment *model.Comment) error
	GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error)
}

type useCase struct {
	commentRepository commentrep.RepositoryI
}

func New(commentRepository commentrep.RepositoryI) UseCaseI {
	return &useCase{
		commentRepository: commentRepository,
	}
}

func (uc *useCase) CreateComment(ctx context.Context, comment *model.Comment) error {
	err := uc.commentRepository.CreateComment(ctx, comment)
	// if err == nil {
	// 	logrus.Info("Comment succesfully created")
	// }
	return err
}

func (uc *useCase) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
	comments, err := uc.commentRepository.GetComments(ctx, postId, limit, offset)
	// if err == nil {
	// 	logrus.Info("Comment succesfully finded")
	// }
	return comments, err
}
