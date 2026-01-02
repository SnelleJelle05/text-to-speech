import io
import soundfile as sf
from fastapi.responses import StreamingResponse
import outetts
from pydantic import BaseModel

enum = outetts.Models("OuteAI/Llama-OuteTTS-1.0-1B".split("/", 1)[1])
cfg  = outetts.ModelConfig.auto_config(
    enum,
    outetts.Backend.LLAMACPP,
    quantization=outetts.LlamaCppQuantization.Q4_K_M
)
tts  = outetts.Interface(cfg)
speaker = tts.load_default_speaker("EN-FEMALE-1-NEUTRAL")

class TTSRequest(BaseModel):
    text: str

def generate(request: TTSRequest):
    print(request.text)

    # Generate audio but DO NOT call .save()
    result = tts.generate(
        outetts.GenerationConfig(
            text=request.text,
            speaker=speaker,
        )
    )

    # result.audio is a torch tensor: shape [samples] or [1, samples]
    audio = result.audio.squeeze().cpu().numpy()
    sr = result.sr

    # Write WAV
    buffer = io.BytesIO()
    sf.write(buffer, audio, sr, format="WAV", subtype="PCM_16")
    buffer.seek(0)

    return StreamingResponse(
        buffer,
        media_type="audio/wav",
        headers={"Content-Disposition": 'attachment; filename="output.wav"'}
    )
