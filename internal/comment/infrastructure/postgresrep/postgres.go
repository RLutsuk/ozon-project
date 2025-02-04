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

func (dbComment *dataBase) CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	var tx *gorm.DB
	if comment.ParentId != "" {
		tx = dbComment.db.Table("comments").Select("body", "user_id", "post_id", "parent_id", "root_id", "level").Create(comment)
	} else {
		tx = dbComment.db.Table("comments").Select("body", "user_id", "post_id").Create(comment)
	}
	if tx.Error != nil {
		return comment, errors.Wrap(tx.Error, "database error (table comments)")
	}
	return comment, nil
}

func (dbComment *dataBase) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, []*model.Comment, error) {
	rootcomments := make([]*model.Comment, 0, limit)
	var childcomments []*model.Comment
	tx := dbComment.db.Table("comments").Where("post_id = ? AND parent_id IS NULL", postId).Limit(limit).Offset(offset).Order("created_at ASC")
	err := tx.Find(&rootcomments).Error
	if err != nil {
		return nil, nil, err
	}
	tx = dbComment.db.Table("comments").Where("post_id = ? AND parent_id IS NOT NULL", postId).Order("created_at ASC")
	err = tx.Find(&childcomments).Error
	if err != nil {
		return nil, nil, err
	}
	return rootcomments, childcomments, nil
}

func (dbComment *dataBase) GetCommentByiD(ctx context.Context, commentID string) (*model.Comment, error) {
	comment := model.Comment{}

	tx := dbComment.db.Table("comments").Where("id = ?", commentID).Take(&comment)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table comments)")
	}
	return &comment, nil
}
