package postgresrep

import (
	"context"

	"github.com/RLutsuk/ozon-project/graph/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) *dataBase {
	return &dataBase{
		db: db,
	}
}

func (dbPost *dataBase) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	tx := dbPost.db.Table("posts").Select("title", "body", "user_id", "allow_comments").Create(post)
	if tx.Error != nil {
		return post, errors.Wrap(tx.Error, "database error (table posts)")
	}
	return post, nil
}

func (dbPost *dataBase) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	tx := dbPost.db.Table("posts").Where("id = ?", id).Take(&post)
	if tx.Error != nil {
		return &post, errors.Wrap(tx.Error, "database error (table posts)")
	}
	return &post, nil
}

func (dbPost *dataBase) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	tx := dbPost.db.Table("posts")
	err := tx.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
