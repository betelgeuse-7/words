package models

import (
	"fmt"
	"time"
)

type profile struct {
	Id           uint       `json:"id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	RegisteredAt time.Time  `json:"registered_at"`
	Notebooks    []notebook `json:"notebooks"`
}

type definition struct {
	DefinitionId uint      `json:"definition_id"`
	FromLang     language  `json:"from_lang"`
	ToLang       language  `json:"to_lang"`
	Word         string    `json:"word"`
	Meaning      string    `json:"meaning"`
	AddedAt      time.Time `json:"added_at"`
}

type language struct {
	LanguageId   uint   `json:"language_id"`
	Language     string `json:"language"`
	LanguageAbbr string `json:"language_abbr"`
}

type notebook struct {
	NotebookId   uint         `json:"notebook_id"`
	NotebookName string       `json:"notebook_name"`
	IsPublic     bool         `json:"is_public"`
	CreatedAt    time.Time    `json:"created_at"`
	Definitions  []definition `json:"definitions"`
}

/*
type favourite struct {
	Definition definition
}
*/

func (n *notebook) setDefinitions(defs []definition) {
	n.Definitions = defs
}

func (p *profile) setNotebooks(nbs []notebook) {
	p.Notebooks = nbs
}

func (d *definition) setFromLang(l language) {
	d.FromLang = l
}

func (d *definition) setToLang(l language) {
	d.ToLang = l
}

func GetProfile(userId int) (profile, error) {
	var p profile
	userQuery := `select user_id, first_name, last_name, email, registered_at from users where user_id = $1`

	row := db.QueryRow(userQuery, userId)
	err := row.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email, &p.RegisteredAt)
	if err != nil {
		return profile{}, err
	}

	notebooks, err := getNotebooks(userId)
	if err != nil {
		return profile{}, err
	}
	p.setNotebooks(notebooks)

	return p, nil
}

func getDefinitions(notebookId int) ([]definition, error) {
	var defs []definition

	query := `select definition_id, new_word, from_lang, to_lang, meaning, added_at from definitions where notebook = $1`

	rows, err := db.Query(query, notebookId)
	if err != nil {
		return []definition{}, err
	}

	for rows.Next() {
		var def definition
		var fromLangId, toLangId int
		rows.Scan(&def.DefinitionId, &def.Word, &fromLangId, &toLangId, &def.Meaning, &def.AddedAt)

		def.setFromLang(getDefinitionSourceLanguage(fromLangId))
		def.setToLang(getDefinitionTargetLanguage(toLangId))

		defs = append(defs, def)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("getDefinitions: ", err)
		return []definition{}, err
	}

	return defs, nil
}

func getNotebooks(userId int) ([]notebook, error) {
	var notebooks []notebook
	query := `select notebook_id, notebook_name, is_public, created_at from notebooks where owner = $1`

	rows, err := db.Query(query, userId)
	if err != nil {
		return []notebook{}, err
	}

	for rows.Next() {
		var nb notebook
		rows.Scan(&nb.NotebookId, &nb.NotebookName, &nb.IsPublic, &nb.CreatedAt)
		nbDefs, err := getDefinitions(int(nb.NotebookId))
		if err != nil {
			fmt.Println("getNotebooks err: ", err)
			return []notebook{}, nil
		}
		nb.setDefinitions(nbDefs)
		notebooks = append(notebooks, nb)
	}
	if err := rows.Err(); err != nil {
		return []notebook{}, err
	}

	return notebooks, nil
}

func getDefinitionSourceLanguage(fromLangId int) language {
	query := `select u_lang_id, lang, lang_abbr from user_languages where u_lang_id = $1`
	var fromLangInfo language

	row := db.QueryRow(query, fromLangId)
	err := row.Scan(&fromLangInfo.LanguageId, &fromLangInfo.Language, &fromLangInfo.LanguageAbbr)
	if err != nil {
		// ** be sure that err is NOT nil when using Error() method.
		fmt.Println(err)
		return language{}
	}
	return fromLangInfo
}

func getDefinitionTargetLanguage(toLangId int) language {
	query := `select u_lang_id, lang, lang_abbr from user_languages where u_lang_id = $1`
	var toLangInfo language

	row := db.QueryRow(query, toLangId)
	err := row.Scan(&toLangInfo.LanguageId, &toLangInfo.Language, &toLangInfo.LanguageAbbr)
	if err != nil {
		fmt.Println(err)
		return language{}
	}

	return toLangInfo
}
