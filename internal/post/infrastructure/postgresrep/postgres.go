package postgresrep

import (
	"context"
	"fmt"
	"time"

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

func (dbPost *dataBase) CreatePost(ctx context.Context, post *model.Post) error {
	tx := dbPost.db.Table("posts").Create(post)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table posts)")
	}
	return nil
}

func (dbPost *dataBase) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	tx := dbPost.db.Model(&model.Post{}).Table("posts").Where("id = ?", id).Select("id", "title", "body", "user_id", "allow_comments", "created_at").Take(&post)
	if tx.Error != nil {
		return &post, errors.Wrap(tx.Error, "database error (table posts)")
	}

	var date time.Time
	// _ = dbPost.db.Table("posts").Where("id = ?", id).Select("created_at").Take(&post.Created)
	fmt.Println(date)
	fmt.Println(post)
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
