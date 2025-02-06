package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/RLutsuk/ozon-project/graph/model"
	commentrep "github.com/RLutsuk/ozon-project/internal/comment/repository"
	postrep "github.com/RLutsuk/ozon-project/internal/post/repository"
	userrep "github.com/RLutsuk/ozon-project/internal/user/repository"
)

// type BaseDBClient interface {
// 	Get(interface{}, string, ...interface{}) error
// }

// type TestDBClient struct {
// 	success bool
// }

// func (tс *TestDBClient) Get(interface{}, string, ...interface{}) error {
// 	if tс.success {
// 		return nil
// 	}
// 	return fmt.Errorf("This is a test error")
// }

type PostDBClient struct {
	success bool
}

func (p *PostDBClient) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	if p.success {
		return &model.Post{ID: "1", Title: post.Title, Body: post.Body, UserID: post.UserID, Allowcomments: post.Allowcomments, Created: time.Time{}}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

func (p *PostDBClient) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	if p.success {
		return &model.Post{ID: "1", Title: "1", Body: "body post", UserID: "1", Allowcomments: true, Created: time.Time{}}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

func (p *PostDBClient) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	if p.success {
		return []*model.Post{{ID: "1", Title: "Test Post"}}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

type UserDBClient struct {
	success bool
}

func (p *UserDBClient) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	if p.success {
		return &model.User{ID: id}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

func (p *UserDBClient) GetUsers(ctx context.Context, ids []string) ([]*model.User, error) {
	if p.success {
		return []*model.User{{ID: "1"}}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

type CommentDBClient struct {
	success bool
}

func (p *CommentDBClient) CreateComment(ctx context.Context, post *model.Comment) (*model.Comment, error) {
	if p.success {
		return &model.Comment{ID: "1"}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

func (p *CommentDBClient) GetCommentByID(ctx context.Context, id string) (*model.Comment, error) {
	if p.success {
		return &model.Comment{ID: id}, nil
	}
	return nil, fmt.Errorf("This is a test error")
}

func (p *CommentDBClient) GetComments(context.Context, string, int, int) ([]*model.Comment, []*model.Comment, error) {
	if p.success {
		// return []*model.Comment{{ID: "1"}}, []*model.Comment{{ID: "2"}}, nil
		return nil, nil, nil
	}
	return nil, nil, fmt.Errorf("This is a test error")
}

func TestNew(t *testing.T) {
	type args struct {
		postRepository    postrep.RepositoryI
		userRepository    userrep.RepositoryI
		commentRepository commentrep.RepositoryI
	}
	tests := []struct {
		name string
		args args
		want UseCaseI
	}{
		{
			name: "all DBClient exist",
			args: args{
				postRepository:    &PostDBClient{success: true},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
			want: &useCase{
				postRepository:    &PostDBClient{success: true},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
		},
		{
			name: "DBClient not exist",
			args: args{
				postRepository:    &PostDBClient{success: false},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
			want: &useCase{
				postRepository:    &PostDBClient{success: false},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.postRepository, tt.args.userRepository, tt.args.commentRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_CreatePost(t *testing.T) {
	type fields struct {
		postRepository    postrep.RepositoryI
		userRepository    userrep.RepositoryI
		commentRepository commentrep.RepositoryI
	}
	type args struct {
		ctx       context.Context
		inputPost model.CreatePostInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Post
		wantErr bool
	}{
		{
			name: "succes creating",
			fields: fields{
				postRepository:    &PostDBClient{success: true},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
			args: args{
				ctx: context.Background(),
				inputPost: model.CreatePostInput{
					Title:         "1",
					Body:          "body post",
					UserID:        "1",
					AllowComments: true,
				},
			},
			want: &model.Post{
				ID:            "1",
				Title:         "1",
				Body:          "body post",
				UserID:        "1",
				Allowcomments: true,
				Created:       time.Time{},
			},
			wantErr: false,
		},
		{
			name: "empty body and title",
			fields: fields{
				postRepository:    &PostDBClient{success: true},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
			args: args{
				ctx: context.Background(),
				inputPost: model.CreatePostInput{
					Title:         "1",
					Body:          "",
					UserID:        "",
					AllowComments: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "darabase error",
			fields: fields{
				postRepository:    &PostDBClient{success: false},
				userRepository:    &UserDBClient{success: false},
				commentRepository: &CommentDBClient{success: true},
			},
			args: args{
				ctx: context.Background(),
				inputPost: model.CreatePostInput{
					Title:         "1",
					Body:          "title",
					UserID:        "body",
					AllowComments: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "user not found",
			fields: fields{
				postRepository:    &PostDBClient{success: true},
				userRepository:    &UserDBClient{success: false},
				commentRepository: &CommentDBClient{success: true},
			},
			args: args{
				ctx: context.Background(),
				inputPost: model.CreatePostInput{
					Title:         "1",
					Body:          "title",
					UserID:        "body",
					AllowComments: true,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &useCase{
				postRepository:    tt.fields.postRepository,
				userRepository:    tt.fields.userRepository,
				commentRepository: tt.fields.commentRepository,
			}
			got, err := uc.CreatePost(tt.args.ctx, tt.args.inputPost)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.CreatePost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetPost(t *testing.T) {
	var limit int32 = 4
	var ptrLimit *int32 = &limit
	var offset int32 = 4
	var ptroffset *int32 = &offset
	type fields struct {
		postRepository    postrep.RepositoryI
		userRepository    userrep.RepositoryI
		commentRepository commentrep.RepositoryI
	}
	type args struct {
		ctx    context.Context
		id     string
		limit  *int32
		offset *int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Post
		wantErr bool
	}{
		{
			name: "succes find",
			fields: fields{
				postRepository:    &PostDBClient{success: true},
				userRepository:    &UserDBClient{success: true},
				commentRepository: &CommentDBClient{success: true},
			},
			args: args{
				ctx:    context.Background(),
				id:     "1",
				limit:  ptrLimit,
				offset: ptroffset,
			},
			want: &model.Post{
				ID:            "1",
				Title:         "1",
				Body:          "body post",
				UserID:        "1",
				Allowcomments: true,
				Created:       time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &useCase{
				postRepository:    tt.fields.postRepository,
				userRepository:    tt.fields.userRepository,
				commentRepository: tt.fields.commentRepository,
			}
			got, err := uc.GetPost(tt.args.ctx, tt.args.id, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.GetPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetAllPosts(t *testing.T) {
	type fields struct {
		postRepository    postrep.RepositoryI
		userRepository    userrep.RepositoryI
		commentRepository commentrep.RepositoryI
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &useCase{
				postRepository:    tt.fields.postRepository,
				userRepository:    tt.fields.userRepository,
				commentRepository: tt.fields.commentRepository,
			}
			got, err := uc.GetAllPosts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetAllPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.GetAllPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetUserByID(t *testing.T) {
	type fields struct {
		postRepository    postrep.RepositoryI
		userRepository    userrep.RepositoryI
		commentRepository commentrep.RepositoryI
	}
	type args struct {
		ctx context.Context
		obj *model.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &useCase{
				postRepository:    tt.fields.postRepository,
				userRepository:    tt.fields.userRepository,
				commentRepository: tt.fields.commentRepository,
			}
			got, err := uc.GetUserByID(tt.args.ctx, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("useCase.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("useCase.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
