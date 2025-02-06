package loaders

import (
	"context"
	"net/http"
	"time"

	"github.com/RLutsuk/ozon-project/graph/model"
	userrep "github.com/RLutsuk/ozon-project/internal/user/repository"
	"github.com/sirupsen/logrus"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type userReader struct {
	userRepository userrep.RepositoryI
}

func (u *userReader) getUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	users, err := u.userRepository.GetUsers(ctx, userIDs)
	if err != nil {
		logrus.WithError(err).Error("post resolver error: get user")
		return nil, nil
	}
	return users, nil
}

type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

func NewLoaders(userRepository userrep.RepositoryI) *Loaders {
	ur := &userReader{userRepository: userRepository}
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(ur.getUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}
func Middleware(userRepository userrep.RepositoryI, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(userRepository)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
