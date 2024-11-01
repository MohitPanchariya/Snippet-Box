package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

// ER_DUP_ENTRY is the error returned by MySQL on trying to insert
// a duplicate key
var ER_DUP_ENTRY = uint16(1062)

// Inserts a user record in the database
func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := "INSERT INTO users(name, email, hashed_password, created) values(?, ?, ?, UTC_TIMESTAMP())"
	_, err = m.DB.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == ER_DUP_ENTRY && strings.Contains(mySQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

// Checks if a user with given email and password exist
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Checks if a user exists with a given id
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
