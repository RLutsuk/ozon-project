package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen generate

import (
	commentresolver "github.com/RLutsuk/ozon-project/internal/comment/resolver"
	postresolver "github.com/RLutsuk/ozon-project/internal/post/resolver"
)

type Resolver struct {
	PostResolver    *postresolver.PostResolver
	CommentResolver *commentresolver.CommentResolver
}
