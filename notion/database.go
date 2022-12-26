package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const parentPageID = "abe3e39b9a2a4a34a5d2cd954a3a36c7"

type Url struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
type Note struct {
	Type  string  `json:"type"`
	Title []Title `json:"title"`
}

type NoteProperty struct {
	Note Note `json:"Note"`
	Url  Url  `json:"URL"`
}

type DBParent struct {
	Type string `json:"type"`
	ID   string `json:"database_id"`
}

type CreateNote struct {
	Properties NoteProperty `json:"properties"`
	Parent     DBParent     `json:"parent"`
}

func insertNote(noteName, noteUrl string) error {
	title := Title{
		Text: Text{
			Content: noteName,
		},
	}
	note := Note{
		Type: "title",
		Title: []Title{
			title,
		},
	}

	url := Url{
		Type: "url",
		Url:  noteUrl,
	}
	createNote := NoteProperty{
		Note: note,
		Url:  url,
	}

	createPage := CreateNote{
		Properties: createNote,
		Parent: DBParent{
			Type: "database_id",
			ID:   parentPageID,
		},
	}

	b, err := json.Marshal(createPage)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error creating page: %s", responseBody)
	}
	var page Page
	if err := json.Unmarshal(responseBody, &page); err != nil {
		return err
	}
	return nil
}
