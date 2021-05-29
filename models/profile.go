package models

import (
	"fmt"
	"log"
	"time"

	"github.com/betelgeuse-7/words/constants"
)

type profile struct {
	Id                         uint
	FirstName, LastName, Email string
	RegisteredAt               time.Time
	Notebooks                  []notebook
}

type definition struct {
	DefinitionId                           uint
	FromLangId, ToLangId                   uint
	FromLangInformation, ToLangInformation language
	Word, Meaning                          string
	AddedAt                                time.Time
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

// ! !!
func getDefinitions(notebookId int) ([]definition, error) {
	query := `select definition_id, from_lang, new_word, to_lang
			meaning, added_at from definitions where notebook = $1`
	rows, err := db.Query(query, notebookId)
	if err != nil {
		fmt.Println("profile.go#getDefinitions ERR: ", err)
		return []definition{}, err
	}

	var defs []definition

	for rows.Next() {
		var def definition
		rows.Scan(&def.DefinitionId, &def.FromLangId, &def.Word, &def.ToLangId, &def.Meaning, &def.AddedAt)
		def.FromLangInformation, err = getDefinitionSourceLanguage(int(def.FromLangId))
		if err != nil {
			log.Println("models/profile.go#getDefinitions ERR2: ", err)
			return []definition{}, err
		}
		def.ToLangInformation, err = getDefinitionTargetLanguage(int(def.ToLangId))
		if err != nil {
			log.Println("models/profile.go#getDefinitions ERR3: ", err)
		}
		fmt.Println("FROM LANG INFO: ", def.FromLangInformation)
		fmt.Println(def)
		defs = append(defs, def)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return []definition{}, err
	}

	return defs, nil
}

func getNotebooks(userId int) ([]notebook, error) {
	query := `select notebook_id, notebook_name, is_public, created_at from notebooks where owner = $1`
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

func getDefinitionSourceLanguage(fromLangId int) (language, error) {
	query := `select u_lang_id, lang, lang_abbr from user_languages where u_lang_id = $1`
	var fromLangInfo language

	row := db.QueryRow(query, fromLangId)
	if err := row.Scan(&fromLangInfo.LanguageId, &fromLangInfo.Language, &fromLangInfo.LanguageAbbr); err.Error() == constants.SQL_NO_ROWS_ERROR {
		return language{}, nil
	}

	return fromLangInfo, nil
}

func getDefinitionTargetLanguage(toLangId int) (language, error) {
	query := `select u_lang_id, lang, lang_abbr from user_languages where u_lang_id = $1`
	var toLangInfo language

	row := db.QueryRow(query, toLangId)
	if err := row.Scan(&toLangInfo.LanguageId, &toLangInfo.Language, &toLangInfo.LanguageAbbr); err.Error() == constants.SQL_NO_ROWS_ERROR {
		return language{}, nil
	}
	return toLangInfo, nil
}
