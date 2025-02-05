package usecase

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	commentrep "github.com/RLutsuk/ozon-project/internal/comment/repository"
	userrep "github.com/RLutsuk/ozon-project/internal/user/repository"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type UseCaseI interface {
	CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error)
	GetUserByID(ctx context.Context, obj *model.Comment) (*model.User, error)
	// GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error)
}

type useCase struct {
	commentRepository commentrep.RepositoryI
	userRepository    userrep.RepositoryI
}

func New(commentRepository commentrep.RepositoryI, userRepository userrep.RepositoryI) UseCaseI {
	return &useCase{
		commentRepository: commentRepository,
		userRepository:    userRepository,
	}
}

func (uc *useCase) CreateComment(ctx context.Context, input model.CreateCommentInput) (*model.Comment, error) {
	var newcomment model.Comment
	newcomment.Body = input.Body
	newcomment.PostId = input.PostID
	newcomment.UserId = input.UserID
	lenComment := []rune(newcomment.Body)

	if len(lenComment) > 2000 {
		return nil, errors.Wrap(model.ErrBadData, "uc error: comment is too big (method CreateComment)")
	}

	if input.ParentID != "" {
		newcomment.ParentId = input.ParentID
		parentComment, err := uc.commentRepository.GetCommentByiD(ctx, input.ParentID)
		if err != nil {
			return nil, errors.Wrap(err, "uc error: failed to get comment (method CreateComment)")
		}
		newcomment.Level = parentComment.Level + 1
		if newcomment.Level == 1 {
			newcomment.RootID = input.ParentID
		} else {
			newcomment.RootID = parentComment.RootID
		}
	}

	comment, err := uc.commentRepository.CreateComment(ctx, &newcomment)
	if err == nil {
		logrus.Info("Comment succesfully created")
	} else {
		return nil, errors.Wrap(err, "uc error: failed to create comment (method CreateComment)")
	}
	return comment, err
}

func (uc *useCase) GetUserByID(ctx context.Context, obj *model.Comment) (*model.User, error) {
	user, err := uc.userRepository.GetUserByID(ctx, obj.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "uc error: failed to get user by id (method GetUserByID)")
	}
	return user, nil
}

// func (uc *useCase) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
// 	rootcomments, _, err := uc.commentRepository.GetComments(ctx, postId, limit, offset)
// 	if err == nil {
// 		logrus.Info("Comment succesfully found")
// 	}
// 	return rootcomments, err
// }
