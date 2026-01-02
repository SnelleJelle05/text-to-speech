from fastapi import FastAPI
import textToSpeechApi
from pydantic import BaseModel

app = FastAPI()


@app.get("/health")
def health_check():
    return {"status": "ok"}


class TTSRequest(BaseModel):
    text: str


@app.post("/text-to-speech")
def text_to_speech_api(request: TTSRequest):
    return  textToSpeechApi.generate(request)