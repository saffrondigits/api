package repo

import (
	"github.com/saffrondigits/api/models"

	"database/sql"
)

type SqlDbImplementation interface {
	RegisterUser(user models.User) error
}

type sqlDBQuery struct {
	dbConn *sql.DB
}

func NewSqlDbImplementation(db *sql.DB) SqlDbImplementation {
	return &sqlDBQuery{
		dbConn: db,
	}
}

func (db *sqlDBQuery) RegisterUser(user models.User) error {
	query := `INSERT INTO user_account (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := db.dbConn.Exec(query, user.FirstName, user.LastName, user.Email, user.Password)
	return err
}
