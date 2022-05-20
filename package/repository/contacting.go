package repository

import (
	"fmt"
	"norbekov/model"
	"norbekov/util/logrus"

	"github.com/jmoiron/sqlx"
)

type ContactPostDB struct {
	db *sqlx.DB
}

func NewContactPostDB(db *sqlx.DB) *ContactPostDB {
	return &ContactPostDB{db: db}
}

func (repo *ContactPostDB) CreateContactPost(post model.Contact, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name ,last_name  , phone_number , type_service,text) VALUES ($1, $2, $3,$4,$5)  RETURNING id", contact)

	row := repo.db.QueryRow(query, post.FirstName, post.LastName, post.PhoneNumber, post.TypeService, post.Text)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}
