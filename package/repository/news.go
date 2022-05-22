package repository

import (
	// "fmt"
	// "norbekov/model"
	// "norbekov/util/logrus"

	"fmt"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/util/logrus"
	"time"

	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
	// "github.com/lib/pq"
)

type NewsPostDB struct {
	db *sqlx.DB
}

func NewNewsPostDB(db *sqlx.DB) *NewsPostDB {
	return &NewsPostDB{db: db}
}

func (repo *NewsPostDB) CreateNewsPost(post model.NewsPost, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_title ,post_img_url,   post_body ) VALUES ($1, $2, $3)  RETURNING id", news)

	row := repo.db.QueryRow(query, post.PostTitle, post.PostImgUrl, post.PostBody)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *NewsPostDB) UpdateNewsImage(ID int, filePath string, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("UPDATE %s  SET post_img_path = $1,updated_at = $2    WHERE id = $3 RETURNING id", news)

	rows, err := repo.db.Exec(query, filePath, tm, ID)

	if err != nil {
		logrus.Errorf("ERROR: Update PostImage failed : %v", err)
		return 0, err
	}

	effectedRowsNum, err := rows.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Update Post Image effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update  image")
	return effectedRowsNum, nil

}

func (repo *NewsPostDB) UpdateNews(Id int, post model.NewsPost, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("	UPDATE %s SET post_title =$1, post_img_url  = $2, post_body = $3,  updated_at=$4 WHERE id = $5 RETURNING id", news)
	rows, err := repo.db.Exec(query, post.PostTitle, post.PostImgUrl, post.PostBody, tm, Id)

	if err != nil {
		logrus.Errorf("ERROR: Update home : %v", err)
		return 0, err
	}
	effectedRowsNum, err := rows.RowsAffected()
	if err != nil {
		logrus.Errorf("ERROR: Update Home effectedRowsNum failed : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update l")
	return effectedRowsNum, nil
}
func (repo *NewsPostDB) GetNewsById(id string, logrus *logrus.Logger) (model.NewsFull, error) {

	var post model.NewsFull
	query := fmt.Sprintf("SELECT  id, post_title, post_img_path,post_img_url, post_body, post_date  FROM %s WHERE id=$1 ", news)
	err := repo.db.Get(&post, query, id)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return post, err
	}
	logrus.Info("DONE:get user data from psql")

	return post, err
}

func (repo *NewsPostDB) DeleteNews(id string, logrus *logrus.Logger) error {

	_, err := repo.db.Exec("DELETE from news WHERE id = $1", id)
	if err != nil {
		logrus.Errorf("ERROR: Update news : %v", err)
		return err
	}
	return nil
}
