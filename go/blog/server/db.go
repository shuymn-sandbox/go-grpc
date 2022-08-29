package main

import (
	"sync"
	"time"
)

type PostStatus int

const (
	unspecified PostStatus = iota
	draft
	private
	unlisted
	public
)

type Post struct {
	ID          int
	AuthorID    int
	Title       string
	Content     []byte
	Status      PostStatus
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID   int
	Name string
}

type DB interface {
	GetPost(id int) (*Post, bool)
	GetUser(id int) (*User, bool)
}

type db struct {
	posts map[int]*Post
	users map[int]*User
	sync.RWMutex
}

func NewDB() DB {
	return &db{
		posts: map[int]*Post{
			1: {
				ID:       1,
				AuthorID: 1,
				Title:    "Hello",
				Content: []byte(`# Title
sample text. sample text.  
sample text. sample text.  
`),
				Status:      public,
				PublishedAt: time.Now(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
		users: map[int]*User{
			1: {
				ID:   1,
				Name: "one",
			},
		},
	}
}

func (d *db) GetPost(id int) (_ *Post, ok bool) {
	d.RLock()
	defer d.RUnlock()
	post, ok := d.posts[id]
	return post, ok
}

func (d *db) GetUser(id int) (_ *User, ok bool) {
	d.RLock()
	defer d.RUnlock()
	user, ok := d.users[id]
	return user, ok
}
