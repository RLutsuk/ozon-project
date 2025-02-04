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
	if tx.Error != nil {
		return &user, errors.Wrap(tx.Error, "database error (table users)")
	}
	return &user, nil
}
