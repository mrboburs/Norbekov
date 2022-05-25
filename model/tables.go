package model

type TablePost struct {
	ID        int    `json:"-"`
	PostTitle string `json:"post_title"`

	PostImgUrl string `json:"post_img_url"`
	PostBody   string `json:"post_body"`
	Price      string `json:"price" `
	Duration   string `json:"duration" `
}

type TableFull struct {
	ID          int    `json:"id" db:"id"`
	PostTitle   string `json:"post_title" db:"post_title"`
	PostImgPath string `json:"post_img_path" db:"post_img_path"`
	PostImgUrl  string `json:"post_img_url" db:"post_img_url"`
	PostBody    string `json:"post_body" db:"post_body"`
	PostDate    string `json:"post_date" db:"post_date"`
	Price       string `json:"price" db:"price"`
	Duration    string `json:"duration" db:"duration"`
}
type allTable struct {
	AllHome []TableFull
}

type CourseFull struct {
	Title    string `json:"post_title" `
	Body     string `json:"post_body" `
	Price    string `json:"price" `
	Duration string `json:"duration" `
	Term     string `json:"term" `
	Format   string `json:"format" `
	Date     string `json:"date" `
}

type CourseFull1 struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"post_title" db:"post_title"`
	Body     string `json:"post_body" db:"post_body"`
	Price    string `json:"price" db:"price"`
	Duration string `json:"duration" db:"duration"`
	Term     string `json:"term" db:"term"`
	Format   string `json:"format" db:"format"`
	Date     string `json:"date" db:"created_at"`
}
