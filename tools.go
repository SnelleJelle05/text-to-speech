package textToSpeech

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type LanguageDetectionData struct {
	Language string   `json:"language"`
	Words    []string `json:"words"`
}

// ISO 639-1
// English is defined as "en"
// https://en.wikipedia.org/wiki/ISO_639-1
func detectLanguage(text string) string {
	language := "en"
	words := strings.Split(text, " ")

	dirs, _ := os.ReadDir("languages")
	for _, dir := range dirs {
		file, err := os.Open("languages/" + dir.Name())
		if err != nil {
			panic(err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		var jsonData LanguageDetectionData
		err = decoder.Decode(&jsonData)
		if err != nil {
			panic(err)
		}

		for _, word := range words {
			word = removeSpecialCharacters(word)
			for _, jsonWord := range jsonData.Words {
				if strings.ToLower(jsonWord) == strings.ToLower(word) {
					language = jsonData.Language
				}
			}
		}
	}

	return language
}

func removeSpecialCharacters(str string) string {
	re, err := regexp.Compile("[^\\p{L}\\p{N}]+")
	if err != nil {
		panic(err)
	}

	return re.ReplaceAllString(str, "")
}

func nameGenerator(language string) string {
	currentTime := time.Now().UnixNano()

	return fmt.Sprintf("%v_%v", language, currentTime)
}
