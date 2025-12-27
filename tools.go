package textToSpeech

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"strings"
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

	globalSameWordCounter := 0

	dirs, _ := os.ReadDir("languages")
	for _, dir := range dirs {
		localSameWordCounter := 0

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
					localSameWordCounter++
				}
			}
		}

		// Check if the local counter is higher than the global counter
		// this ensures we get the language with the most matching words
		if localSameWordCounter > globalSameWordCounter {
			globalSameWordCounter = localSameWordCounter
			language = jsonData.Language
		}

	}

	return language
}

func writeMp3File(text string, language string) io.Reader {
	if len(text) > 200 {
		panic("Text length exceeds the maximum 200 characters")
	}

	url := fmt.Sprintf("http://translate.google.com/translate_tts?client=tw-ob&tl=%s&q=%s&textlen=%d", language, url2.QueryEscape(text), len(text))
	responseMp3File, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	return responseMp3File.Body
}

func removeSpecialCharacters(str string) string {
	re, err := regexp.Compile("[^\\p{L}\\p{N}]+")
	if err != nil {
		panic(err)
	}

	return re.ReplaceAllString(str, "")
}
