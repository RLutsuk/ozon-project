package model

import "github.com/pkg/errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrBadData            = errors.New("bad data")
	ErrPostNotFound       = errors.New("post not found")
	ErrCommentNotFound    = errors.New("comment not found")
	ErrСommentsProhibited = errors.New("comments on this post are prohibited")
	ErrPostsDontExist     = errors.New("posts don't exist")
	// ErrCommentsDontExist  = errors.New("comments don't exist")
	ErrCommentOffset = errors.New("offset exceeds number of comments")
	// ErrDataBase           = errors.New("database error")
)
