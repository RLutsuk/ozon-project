package usecase

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	commentrep "github.com/RLutsuk/ozon-project/internal/comment/repository"
	postrep "github.com/RLutsuk/ozon-project/internal/post/repository"
	userrep "github.com/RLutsuk/ozon-project/internal/user/repository"
	"github.com/sirupsen/logrus"
)

type UseCaseI interface {
	CreatePost(ctx context.Context, inputPost model.CreatePostInput) (*model.Post, error)
	GetPost(ctx context.Context, id string, limit, offset *int32) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
	GetUserByID(ctx context.Context, obj *model.Post) (*model.User, error)
}

type useCase struct {
	postRepository    postrep.RepositoryI
	userRepository    userrep.RepositoryI
	commentRepository commentrep.RepositoryI
}

func New(postRepository postrep.RepositoryI, userRepository userrep.RepositoryI, commentRepository commentrep.RepositoryI) UseCaseI {
	return &useCase{
		postRepository:    postRepository,
		userRepository:    userRepository,
		commentRepository: commentRepository,
	}
}

func (uc *useCase) CreatePost(ctx context.Context, inputPost model.CreatePostInput) (*model.Post, error) {
	var newpost model.Post
	newpost.Title = inputPost.Title
	newpost.Allowcomments = inputPost.AllowComments
	newpost.Body = inputPost.Body
	newpost.UserID = inputPost.UserID
	post, err := uc.postRepository.CreatePost(ctx, &newpost)
	if err == nil {
		logrus.Info("Post succesfully created")
	}
	return post, err
}

func (uc *useCase) GetPost(ctx context.Context, id string, limit, offset *int32) (*model.Post, error) {
	post, _ := uc.postRepository.GetPostByID(ctx, id)
	limitInt := 0
	offsetInt := 0
	if limit == nil {
		limitInt = 10
	} else {
		limitInt = int(*limit)
	}

	if limitInt == 0 {
		limitInt = 10
	}

	if offset != nil {
		offsetInt = int(*offset)
	}

	rootComments, childcomments, err := uc.commentRepository.GetComments(ctx, post.ID, limitInt, offsetInt)
	if len(childcomments) == 0 {
		post.Comments = rootComments
		return post, nil
	}
	commentMap := make(map[int]map[string][]*model.Comment)
	for _, comment := range childcomments {
		level := comment.Level
		if _, exists := commentMap[level]; !exists {
			commentMap[level] = make(map[string][]*model.Comment)
		}
		commentMap[level][comment.ParentId] = append(commentMap[level][comment.ParentId], comment)
	}

	var buildTree func(parentId string, level int) []*model.Comment
	buildTree = func(parentId string, level int) []*model.Comment {
		comments := commentMap[level][parentId]
		for i := range comments {
			comments[i].Replies = buildTree(comments[i].ID, level+1)
		}
		return comments
	}

	for _, rootComment := range rootComments {
		rootComment.Replies = buildTree(rootComment.ID, 1)
	}

	post.Comments = rootComments
	if err == nil {
		logrus.Info("Post succesfully found")
	}
	return post, err
}

func (uc *useCase) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	posts, err := uc.postRepository.GetAllPosts(ctx)
	// if err == nil {
	// 	logrus.Info("Posts succesfully found")
	// }
	return posts, err
}

func (uc *useCase) GetUserByID(ctx context.Context, obj *model.Post) (*model.User, error) {
	user, err := uc.userRepository.GetUserByID(ctx, obj.UserID)
	return user, err
}
