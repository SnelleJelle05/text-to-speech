package textToSpeech

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
)

// InitSpeaker WAV format is used for Python TTS API
func InitSpeaker() {
	err := speaker.Init(44100, 44100)
	if err != nil {
		return
	}
}

func TextToSpeech(text string) {
	payload := fmt.Sprintf(`{"text": %q}`, text)

	resp, err := http.Post(
		fmt.Sprintf("%s/text-to-speech", pythonApiUrl),
		"application/json",
		strings.NewReader(payload),
	)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		InitPythonApi()
	}
	file := resp.Body

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	streamer, _, err := wav.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	speaker.PlayAndWait(streamer)
}
