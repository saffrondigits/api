package repo

import (
	"errors"

	"github.com/saffrondigits/api/models"
	"github.com/sirupsen/logrus"

	"database/sql"
)

type SqlDbImplementation interface {
	RegisterUser(user models.User) error
	CheckIfEmailExists(email string) (*models.DbLoginCred, error)
	DeleteCheckedUser(email string) error
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

func (db *sqlDBQuery) CheckIfEmailExists(email string) (*models.DbLoginCred, error) {
	var emailId, hash string
	err := db.dbConn.QueryRow("SELECT email, password FROM user_account WHERE email=$1", email).Scan(&emailId, &hash)
	if err != nil {
		logrus.Errorf("error while retrieving the user: %v", err)
		if err == sql.ErrNoRows {
			return nil, errors.New("didn't find any user with this email")
		}
		return nil, err
	}

	return &models.DbLoginCred{Email: emailId, Hash: hash}, nil
}

func (db *sqlDBQuery) DeleteCheckedUser(email string) error {
	row := db.dbConn.QueryRow("DELETE FROM user_account WHERE email=$1", email)
	if row.Err() != nil {
		logrus.Errorf("Error while deleting the user: %v", row.Err())
		return row.Err()
	}
	return nil
}
