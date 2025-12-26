package main

import (
	"fmt"

	"io"
	"net/http"
	url2 "net/url"
	"os"
	"time"
)

func main() {
	createFolderIfNotExist()

	text := "Hello, this is a sample text to speech conversion using Go!"
	speak(text)
}

func createFolderIfNotExist() {
	err := os.Mkdir("speech", 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func speak(text string) {
	fileName := writeMp3File(text)

	fmt.Println("Generated file:", fileName)
}

func writeMp3File(text string) string {
	if len(text) > 200 {
		panic("Text length exceeds the maximum 200 characters")
	}

	language := "en"
	url := fmt.Sprintf("http://translate.google.com/translate_tts?client=tw-ob&tl=%s&q=%s&textlen=%d", language, url2.QueryEscape(text), len(text))
	fmt.Println(url)
	responseMp3File, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(responseMp3File.Body)
	fmt.Println(responseMp3File)

	fileName := nameGenerator(language)
	mp3Path := fmt.Sprintf("speech/%v.mp3", fileName)
	file, err := os.Create(mp3Path)
	if err != nil {
		panic(err)
	}

	// Replaces the content of file with the content of response body || mp3 file
	io.Copy(file, responseMp3File.Body)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return fileName
}

func nameGenerator(language string) string {
	currentTime := time.Now().UnixNano()

	return fmt.Sprintf("%v_%v", language, currentTime)
}
