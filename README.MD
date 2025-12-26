# textToSpeech (Go)

Detect language from simple wordlists, fetch Google Translate TTS MP3, and play it locally.

## Install

```bash
go get <module>/tts
```

## Usage

```go
package main

import "..."

func main() {
	textToSpeech.SpeakInit()
	textToSpeech.Speak("Hello world")
}
```

## Folders

- `./speech/` (auto-created) — downloaded `*.mp3`
- `./languages/` (required) — JSON wordlists for detection

Example `languages/english.json`:

```json
{ "language": "en", "words": ["the", "and", "you", "hello"] }
```

## Notes

- Max input length: **200** chars (panics if exceeded)
- Uses an unofficial Google TTS endpoint; network/rate-limit issues may occur
