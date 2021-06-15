package models

import (
	"log"
	"time"
)

// https://golangbot.com/interfaces-part-1/

type publicUser struct {
	UserId       uint      `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	RegisteredAt time.Time `json:"registered_at"`
}

type userCred struct {
	UserId   uint
	Password string
}

func Register(firstname, lastname, email, password string, registeredAt time.Time) error {
	_, err := db.Exec(`insert into 
				users (first_name, last_name, email, password, registered_at)
				values ($1, $2, $3, $4, $5);`, firstname, lastname, email, password, registeredAt)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]publicUser, error) {
	query := `select user_id, first_name, last_name, registered_at from users;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var us []publicUser

	for rows.Next() {
		var u publicUser
		if err = rows.Scan(&u.UserId, &u.FirstName, &u.LastName, &u.RegisteredAt); err != nil {
			log.Println(err, "GetUsers")
			return nil, err
		}
		us = append(us, u)
	}
	if err := rows.Err(); err != nil {
		log.Println(err, "GetUsers2")
		return nil, err
	}
	return us, nil
}

func GetSingleUser(userId int) (publicUser, error) {
	query := `select user_id, first_name, last_name, registered_at from users where user_id = $1`
	row := db.QueryRow(query, userId)

	var user publicUser

	err := row.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.RegisteredAt)
	if err != nil {
		return publicUser{}, err
	}

	return user, nil
}

func GetUserCredsByEmail(email string) (userCred, error) {
	query := `select user_id, password from users where email=$1`
	row := db.QueryRow(query, email)

	var user userCred

	err := row.Scan(&user.UserId, &user.Password)
	if err != nil {
		return userCred{}, err
	}
	return user, nil
}
