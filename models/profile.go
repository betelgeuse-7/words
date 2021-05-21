package models

import (
	"fmt"
	"time"
)

type profile struct {
	Id                         uint
	FirstName, LastName, Email string
	RegisteredAt               time.Time
	Notebooks                  []notebook
}

type definition struct {
	DefinitionId     uint
	FromLang, ToLang language
	Word, Meaning    string
	Notebook         notebook
	AddedAt          time.Time
}

type language struct {
	LanguageId             uint
	Language, LanguageAbbr string
}

type notebook struct {
	NotebookId   uint
	NotebookName string
	IsPublic     bool
	CreatedAt    time.Time
	Definitions  []definition
}

type favourite struct {
	Definition definition
}

// ! notebooks empty.return defs, nil

func GetProfile(userId int) (profile, error) {
	var p profile
	userQuery := `select user_id, first_name, last_name, email, registered_at from users where user_id = $1`

	row := db.QueryRow(userQuery, userId)
	err := row.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email, &p.RegisteredAt)
	if err != nil {
		fmt.Println(err)
		return profile{}, err
	}

	notebooks, err := getNotebooks(userId)
	if err != nil {
		fmt.Println("GetProfile ERR: ", err)
		return profile{}, err
	}
	p.Notebooks = notebooks

	return p, nil
}

func getDefinitions(notebookId int) ([]definition, error) {
	query := `select * from definitions where notebook = $1`
	rows, err := db.Query(query, notebookId)
	if err != nil {
		fmt.Println("profile.go#getDefinitions ERR: ", err)
		return []definition{}, err
	}

	var defs []definition

	for rows.Next() {
		var def definition
		rows.Scan(&def.DefinitionId, &def.FromLang, &def.Word, &def.ToLang, &def.Meaning, &def.Notebook, &def.AddedAt)
		defs = append(defs, def)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return []definition{}, err
	}

	return defs, nil
}

func getNotebooks(userId int) ([]notebook, error) {
	query := `select * from notebooks where owner = $1`
	rows, err := db.Query(query, userId)
	if err != nil {
		fmt.Println(err)
		return []notebook{}, err
	}

	var notebooks []notebook

	for rows.Next() {
		var nb notebook
		rows.Scan(&nb.NotebookId, &nb.NotebookName, &nb.IsPublic, &nb.CreatedAt)

		nbDefs, err := getDefinitions(int(nb.NotebookId))
		if err != nil {
			fmt.Println(err)
			return []notebook{}, nil
		}
		nb.Definitions = nbDefs
		notebooks = append(notebooks, nb)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return []notebook{}, err
	}

	return notebooks, nil
}
