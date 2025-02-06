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

func (dbUser *dataBase) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	tx := dbUser.db.Table("users").Where("id = ?", id).Take(&user)
	if tx.RowsAffected == 0 {
		return nil, errors.Wrap(model.ErrUserNotFound, "database error: user not found (method GetUserByID, table users)")
	}
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error: internal (method GetUserByID, table users)")
	}
	return &user, nil
}

func (dbUser *dataBase) GetUsers(ctx context.Context, id []string) ([]*model.User, error) {
	var users []*model.User
	tx := dbUser.db.Table("users").Where("id IN ?", id).Find(&users)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error: internal (method GetUsers, table users)")
	}
	return users, nil
}
