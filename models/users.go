package models

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// https://golangbot.com/interfaces-part-1/

type newUser struct {
	FirstName, LastName, Email, Password string
}

type publicUser struct {
	UserId       uint      `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	RegisteredAt time.Time `json:"registered_at"`
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wantHeader := map[string]string{
		"Content-Type": "application/json",
	}
	if ok, msg := utils.CheckRequestHeader(wantHeader, r.Header); ok && msg == "ok" {
		var new newUser
		if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
			json.NewEncoder(w).Encode(responses.REGISTER_FAIL)
			return
		}
		if err := utils.LenGreaterThanZero(new.FirstName, new.LastName, new.Email, new.Password); err != nil {
			json.NewEncoder(w).Encode(responses.MISSING_CREDENTIALS)
			return
		}
		if err := utils.ValidateEmail(new.Email); err != nil {
			json.NewEncoder(w).Encode(responses.EMAIL_INVALID)
			return
		}

		// * encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(responses.SERVER_ERROR)
			return
		}

		new.Password = string(password)
		registeredAt := time.Now()

		_, err = db.Exec(`insert into 
				users (first_name, last_name, email, password, registered_at)
				values ($1, $2, $3, $4, $5);`, new.FirstName, new.LastName, new.Email, new.Password, registeredAt)

		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(responses.SERVER_ERROR)
			return
		}

		// * Success
		json.NewEncoder(w).Encode(responses.REGISTER_SUCCESS)

	} else {
		json.NewEncoder(w).Encode(responses.CHECK_HEADER_FAIL)
		return
	}
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

type userCred struct {
	UserId   uint
	Password string
}

func GetUserCredsByEmail(email string) (userCred, error) {
	query := `select user_id, password from users where email=$1`
	row := db.QueryRow(query, email)

	var user userCred

	err := row.Scan(&user)
	if err != nil {
		return userCred{}, err
	}
	return user, nil
}
