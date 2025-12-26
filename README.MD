# textToSpeech (Go)

Detect language from wordlists, fetch Google Translate TTS MP3, play locally.

## Install
```bash
go get <module>/tts
```

## Use
```go
package main

import textToSpeech "<module>/tts"

func main() {
	textToSpeech.SpeakInit()
	textToSpeech.Speak("Hello world")
}
```

## Needs
- `./speech/` (auto-created)
- `./languages/` (required) JSON wordlists, e.g.
```json
{ "language": "en", "words": ["the", "and", "you", "hello"] }
```
