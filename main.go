package textToSpeech

import (
	"io"

	"time"

	"github.com/gopxl/beep/v2/mp3"
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

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}

	speaker.PlayAndWait(streamer)
}
