package textToSpeech

import (
	"fmt"

	"io"
	"net/http"
	url2 "net/url"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
)

func speakInit() {
	createFolderIfNotExist("speech")
}

func speak(text string) {
	language := detectLanguage(text)
	fileName := writeMp3File(text, language)

	file, err := os.Open(fmt.Sprintf("speech/%v.mp3", fileName))

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	speaker.Init(beep.SampleRate(format.SampleRate), format.SampleRate.N(time.Second/10))
	speaker.PlayAndWait(streamer)
	if err != nil {
		panic(err)
	}
}

func writeMp3File(text string, language string) string {
	if len(text) > 200 {
		panic("Text length exceeds the maximum 200 characters")
	}

	url := fmt.Sprintf("http://translate.google.com/translate_tts?client=tw-ob&tl=%s&q=%s&textlen=%d", language, url2.QueryEscape(text), len(text))
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

	fileName := nameGenerator(language)
	mp3Path := fmt.Sprintf("speech/%v.mp3", fileName)
	file, err := os.Create(mp3Path)
	if err != nil {
		panic(err)
	}

	// Replaces the content of file with the content of response body || mp3 file
	_, err = io.Copy(file, responseMp3File.Body)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return fileName
}
