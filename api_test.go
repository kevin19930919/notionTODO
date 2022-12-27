package webhook

import (
	"testing"
)

// func Test_createNote(t *testing.T) {
// 	if page, err := createNote("test_title", "https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html"); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(page)
// 	}
// }

func Test_getPage(t *testing.T) {
	getPage("68d9990e272c4d9db171a818b937f2a1")
}

// func Test_insertNote(t *testing.T) {
// 	testName := "test_url"
// 	testUrl := "https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html"
// 	if page, err := insertNote(testName, testUrl); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(page)
// 	}
// }
