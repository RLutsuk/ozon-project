package model

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"id"`
	Username  string    `json:"username" gorm:"username"`
	Email     string    `json:"email" gorm:"email"`
	Firstname string    `json:"firstname" gorm:"first_name"`
	Lastname  string    `json:"lastname" gorm:"second_name"`
	Created   time.Time `json:"created" gorm:"column:created_at"`
}

type Post struct {
	ID            string     `json:"id" gorm:"id"`
	Title         string     `json:"title" gorm:"title"`
	Body          string     `json:"body" gorm:"body"`
	User          *User      `json:"user" gorm:"-"`
	UserID        string     `json:"userid" gorm:"user_id"`
	Allowcomments bool       `json:"allowcomments" gorm:"alow_comments"`
	Comments      []*Comment `json:"comments,omitempty" gorm:"-"`
	Created       time.Time  `json:"created" gorm:"column:created_at"`
}

type Comment struct {
	ID       string     `json:"id" gorm:"id"`
	Body     string     `json:"body" gorm:"body"`
	User     *User      `json:"user" gorm:"-"`
	UserId   string     `json:"userid" gorm:"user_id"`
	Post     *Post      `json:"post" gorm:"-"`
	PostId   string     `json:"postid" gorm:"post_id"`
	Created  time.Time  `json:"created" gorm:"column:created_at"`
	Parent   *Comment   `json:"parent,omitempty" gorm:"-"`
	ParentId string     `json:"parentid,omitempty" gorm:"parent_id"`
	Replies  []*Comment `json:"replies,omitempty" gorm:"-"`
}

// func MarshalTimestamp(t time.Time) graphql.Marshaler {
// 	timestamp := t.Unix() * 1000

// 	return graphql.WriterFunc(func(w io.Writer) {
// 		io.WriteString(w, strconv.FormatInt(timestamp, 10))
// 	})
// }

// func UnmarshalTimestamp(v interface{}) (time.Time, error) {
// 	if tmpStr, ok := v.(int); ok {
// 		return time.Unix(int64(tmpStr), 0), nil
// 	}
// 	return time.Time{}, errors.TimeStampError
// }
