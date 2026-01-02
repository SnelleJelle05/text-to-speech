# text-to-speech (Go + local Python TTS)

Minimal Go wrapper that starts/uses a local Python TTS API and plays WAV output locally.

Note: rename to `README.md` (lowercase) if GoLand's Markdown preview doesn't appear.

## Install
Get the package:
```bash
go get github.com/snellejelle05/text-to-speech@latest
```

## Quick start (Go)
```go
import "github.com/snellejelle05/text-to-speech"

func main() {
textToSpeech.InitPythonApi()   // starts local Python API if needed
textToSpeech.InitSpeaker()     // init audio speaker
textToSpeech.TextToSpeech("Hello from Go + local Python TTS")
textToSpeech.StopPythonApi()   // stop the Python process when done
}
```

## Project layout
- /python — Python API, app entry and requirements (expects a .venv and uvicorn)
- Go files — wrapper that starts the Python API, POSTs text, decodes WAV and plays audio.

## Notes
- The Python service must have dependencies installed in `python/.venv`. The Go code attempts to run pip/uvicorn under `../python/.venv/bin/`.
- Audio is expected as WAV; speaker uses sample rate from the response.
- If preview in GoLand fails, rename `/Users/jelle/GolandProjects/tts/README.MD` → `README.md`.

## Troubleshooting
- If requests fail, ensure the Python API is reachable at http://0.0.0.0:8000 and that the virtualenv exists.
- Check stdout/stderr from the Python process for errors when the Go wrapper starts it.
