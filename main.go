package textToSpeech

import (
	"fmt"
	"io"

	"net/http"
	url2 "net/url"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
)

func Speak(text string) {
	language := detectLanguage(text)
	file := writeMp3File(text, language)

	fileReader, ok := file.(io.ReadCloser)
	if !ok {
		fileReader = io.NopCloser(file)
	}

	streamer, format, err := mp3.Decode(fileReader)
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
