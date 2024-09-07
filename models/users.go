package models

import "github.com/mostafejur21/go_rest_api/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return nil
	}

	userID, err := result.LastInsertId()
	u.ID = userID
	return err
}
