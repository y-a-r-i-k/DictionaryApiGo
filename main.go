package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WordType struct {
	
	Word 		string `json:"word"`
	Phonetic 	string `json:"phonetic"`
	Origin 		string `json:"origin"`
	
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`

		Definitions []struct {
			Definition 	string `json:"definition"`
			Example 	string `json:"example"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func main() {

	var word string = os.Args[1]

	var wt []WordType

	var url = "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	var resp, errGet = http.Get(url)

	if (errGet != nil) {
		fmt.Print(errGet)
		return
	}
	defer resp.Body.Close()

	
	var body, err = io.ReadAll(resp.Body)
	if (err != nil) {
		fmt.Print(err)
		return
	}

	var errRead = json.Unmarshal(body, &wt)
	if (errRead != nil) {
		fmt.Print(errRead)
		return
	}

	jsonBytes, err := json.MarshalIndent(wt, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonBytes))
}
