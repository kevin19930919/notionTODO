package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Text struct {
	Content string `json:"content"`
}

type Title struct {
	Text Text `json:"text"`
}

type Desc struct {
	RichText []Title `json:"rich_text"`
}

type Property struct {
	Titles      []Title `json:"title"`
	Description Desc    `json:"Description"`
}
type Parent struct {
	Type string `json:"type"`
	ID   string `json:"database_id"`
}

type CreatePage struct {
	Properties Property `json:"properties"`
	Parent     Parent   `json:"parent"`
	Children   []string `json:"children"`
}

type Page struct {
	ID string `json:"id"`
}

func createNote(titleContent, description string) (*Page, error) {
	title := Title{
		Text: Text{
			Content: titleContent,
		},
	}
	desc := Title{
		Text: Text{
			Content: description,
		},
	}
	properties := Property{
		Titles: []Title{
			title,
		},
		Description: Desc{
			RichText: []Title{
				desc,
			},
		},
	}
	children := []string{}

	createPage := CreatePage{
		Properties: properties,
		Parent: Parent{
			Type: "page_id",
			ID:   parentPageID,
		},
		Children: children,
	}
	b, err := json.Marshal(createPage)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")
	fmt.Println(req)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating page: %s", responseBody)
	}
	var page Page
	if err := json.Unmarshal(responseBody, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func getPage(pageID string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.notion.com/v1/pages/%s", pageID), bytes.NewBuffer([]byte{}))
	if err != nil {
		// return nil, err
		fmt.Println(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")
	fmt.Println(req)
	client := &http.Client{}
	res, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// if res.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("error creating page: %s", responseBody)
	// }
	fmt.Println("response", string(responseBody))
}
