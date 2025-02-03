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

func (dbComment *dataBase) CreateComment(ctx context.Context, comment *model.Comment) error {
	tx := dbComment.db.Table("commets").Create(comment)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table comment)")
	}
	return nil
}

func (dbComment *dataBase) GetComments(ctx context.Context, postId string, limit, offset int) ([]*model.Comment, error) {
	commets := make([]*model.Comment, limit)
	tx := dbComment.db.Table("comments")
	//tx = tx.Limit(limit).Offset(offset).Order("title ASC")
	err := tx.Find(&commets).Error
	if err != nil {
		return nil, err
	}
	return commets, nil
}
