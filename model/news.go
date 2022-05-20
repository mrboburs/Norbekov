package model

import (
// "database/sql"
// "github.com/lib/pq"
)

type NewsPost struct {
	ID        int    `json:"-"`
	PostTitle string `json:"post_title"`

	PostImgUrl string `json:"post_img_url"`
	PostBody   string `json:"post_body"`
}

type NewsFull struct {
	ID          int    `json:"id" db:"id"`
	PostTitle   string `json:"post_title" db:"post_title"`
	PostImgPath string `json:"post_img_path" db:"post_img_path"`
	PostImgUrl  string `json:"post_img_url" db:"post_img_url"`
	PostBody    string `json:"post_body" db:"post_body"`
	PostDate    string `json:"post_date" db:"post_date"`
}
