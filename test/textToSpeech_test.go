package test

import (
	"testing"

	"github.com/snellejelle05/text-to-speech"
)

func TestSpeakEn(t *testing.T) {
	textToSpeech.InitPythonApi()
	textToSpeech.InitSpeaker()
	textToSpeech.TextToSpeech("Hello, how are you?")

	textToSpeech.TextToSpeech("This is a test of the text to speech system.")

	textToSpeech.StopPythonApi()
}
