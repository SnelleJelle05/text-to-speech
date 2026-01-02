package textToSpeech

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const pythonApiUrl = "http://0.0.0.0:8000"

var pythonCmd *exec.Cmd

func InitPythonApi() {
	// Check if the Python API is running by sending a request to the health endpoint
	// maximum wait time 2 seconds
	client := http.Client{Timeout: time.Second * 2}
	resp, err := client.Get(fmt.Sprintf("%s/health", pythonApiUrl))
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil || resp == nil || resp.StatusCode != 200 {
		startPythonApi()
	}
	// TODO : download models if not present
}

func startPythonApi() {
	//Install Python dependencies
	install := exec.Command(
		"../python/.venv/bin/pip",
		"install",
		"-r",
		"requirements.txt",
	)
	install.Dir = "../python"
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr

	if err := install.Run(); err != nil {
		panic(fmt.Errorf("failed to install Python packages: %w", err))
	}

	//Start uvicorn
	cmd := exec.Command(
		"../python/.venv/bin/uvicorn",
		"app:app",
		"--host", "0.0.0.0",
		"--port", "8000",
	)
	cmd.Dir = "../python"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	pythonCmd = cmd
	fmt.Println("Python TTS API started")
}

func StopPythonApi() {
	if pythonCmd != nil && pythonCmd.Process != nil {
		if err := pythonCmd.Process.Kill(); err != nil {
			fmt.Println("Error stopping Python TTS API:", err)
		} else {
			fmt.Println("Python TTS API stopped")
		}
	}
}
