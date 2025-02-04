package model

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"column:username"`
	Email     string    `json:"email" gorm:"column:email"`
	Firstname string    `json:"firstname" gorm:"column:first_name"`
	Lastname  string    `json:"lastname" gorm:"column:second_name"`
	Created   time.Time `json:"created" gorm:"column:created_at"`
}

type Post struct {
	ID            string     `json:"id" gorm:"id;type:uuid;default:uuid_generate_v4()"`
	Title         string     `json:"title" gorm:"column:title"`
	Body          string     `json:"body" gorm:"column:body"`
	User          *User      `json:"user" gorm:"-"`
	UserID        string     `json:"userid" gorm:"column:user_id"`
	Allowcomments bool       `json:"allowcomments" gorm:"column:allow_comments"`
	Comments      []*Comment `json:"comments,omitempty" gorm:"-"`
	Created       time.Time  `json:"created" gorm:"column:created_at"`
}

type Comment struct {
	ID       string     `json:"id" gorm:"id;type:uuid;default:uuid_generate_v4()"`
	Body     string     `json:"body" gorm:"column:body"`
	User     *User      `json:"user" gorm:"-"`
	UserId   string     `json:"userid" gorm:"column:user_id"`
	Post     *Post      `json:"post" gorm:"-"`
	PostId   string     `json:"postid" gorm:"column:post_id"`
	Created  time.Time  `json:"created" gorm:"column:created_at"`
	Parent   *Comment   `json:"parent,omitempty" gorm:"-"`
	ParentId string     `json:"parentid,omitempty" gorm:"column:parent_id"`
	Level    int        `json:"level" gorm:"column:level"`
	RootID   string     `json:"root_comment_id" gorm:"column:root_id"`
	Replies  []*Comment `json:"replies,omitempty" gorm:"-"`
}

type CreateCommentInput struct {
	Body     string `json:"body"`
	UserID   string `json:"userId"`
	PostID   string `json:"postId"`
	ParentID string `json:"parentId,omitempty"`
}
