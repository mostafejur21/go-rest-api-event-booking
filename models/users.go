package models

import (
	"errors"

	"github.com/mostafejur21/go_rest_api/db"
	"github.com/mostafejur21/go_rest_api/utils"
)

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
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return nil
	}

	userID, err := result.LastInsertId()
	u.ID = userID
	return err
}

func (u User) ValidateCredential() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrivePassword string
    err := row.Scan(&retrivePassword)
    if err != nil{

		return errors.New("invalid credentials")
    }
	passwordValid := utils.CheckHashedPassword(u.Password, retrivePassword)

	if !passwordValid {
		return errors.New("invalid credentials")
	}
	return nil
}
