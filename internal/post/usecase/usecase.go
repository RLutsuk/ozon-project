package usecase

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	postrep "github.com/RLutsuk/ozon-project/internal/post/repository"
	"github.com/sirupsen/logrus"
)

type UseCaseI interface {
	CreatePost(ctx context.Context, inputPost model.CreatePostInput) (*model.Post, error)
	GetPost(ctx context.Context, id string) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
}

type useCase struct {
	postRepository postrep.RepositoryI
}

func New(postRepository postrep.RepositoryI) UseCaseI {
	return &useCase{
		postRepository: postRepository,
	}
}

func (uc *useCase) CreatePost(ctx context.Context, inputPost model.CreatePostInput) (*model.Post, error) {
	var newpost model.Post
	newpost.Title = inputPost.Title
	newpost.Allowcomments = inputPost.AllowComments
	newpost.Body = inputPost.Body
	newpost.User.ID = inputPost.UserID
	err := uc.postRepository.CreatePost(ctx, &newpost)
	if err == nil {
		logrus.Info("Post succesfully created")
	}
	post, err := uc.postRepository.GetPostByID(ctx, newpost.ID)
	return post, err
}

func (uc *useCase) GetPost(ctx context.Context, id string) (*model.Post, error) {
	post, err := uc.postRepository.GetPostByID(ctx, id)
	// if err == nil {
	// 	logrus.Info("Post succesfully finded")
	// }
	return post, err
}

func (uc *useCase) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	posts, err := uc.postRepository.GetAllPosts(ctx)
	// if err == nil {
	// 	logrus.Info("Post succesfully finded")
	// }
	return posts, err
}
