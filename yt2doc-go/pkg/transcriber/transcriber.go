package transcriber

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/samar-108/yt2doc-go/internal/models"
)

type whisperOutput struct {
	Segments []struct {
		Start float64 `json:"start"`
		End   float64 `json:"end"`
		Text  string  `json:"text"`
	} `json:"segments"`
}

// convertToWav uses ffmpeg to normalize the audio for Whisper standard (16kHz, Mono).
func convertToWav(ctx context.Context, inputPath string) (string, error) {
	// Create output filename by replacing the extension with .wav
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".wav"

	// Command: ffmpeg -i [input] -ar 16000 -ac 1 -y -acodec pcm_s16le [output]
	cmd := exec.CommandContext(ctx, "ffmpeg",
		"-i", inputPath,
		"-ar", "16000",
		"-ac", "1",
		"-y",
		"-acodec", "pcm_s16le",
		outputPath,
	)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("ffmpeg conversion failed: %w", err)
	}

	return outputPath, nil
}

// Transcribe handles the full transcription pipeline: normalization -> whisper -> parsing.
func Transcribe(ctx context.Context, audioPath string) ([]models.Segment, error) {
	// 1. Normalize audio to WAV
	wavPath, err := convertToWav(ctx, audioPath)
	if err != nil {
		return nil, err
	}
	defer os.Remove(wavPath) // Cleanup WAV but not the original M4A

	// 2. Create a temporary folder for Whisper's JSON output
	tmpDir, err := os.MkdirTemp("", "whisper-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// 3. Run Whisper CLI
	// We use the base model as a good balance of speed and accuracy.
	cmd := exec.CommandContext(ctx, "whisper",
		wavPath,
		"--output_format", "json",
		"--output_dir", tmpDir,
		"--model", "base",
	)

	// Capture stderr to help debug if whisper crashes
	var stderr strings.Builder
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("whisper failed: %v (stderr: %s)", err, stderr.String())
	}

	// 4. Read the generated JSON file
	// Whisper creates a file with the same base name as the input file
	jsonPath := filepath.Join(tmpDir, strings.TrimSuffix(filepath.Base(wavPath), filepath.Ext(wavPath))+".json")
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read whisper output: %w", jsonPath, err)
	}

	// 5. Parse the JSON results
	var output whisperOutput
	if err := json.Unmarshal(jsonData, &output); err != nil {
		return nil, fmt.Errorf("failed to parse transcription results: %w", err)
	}

	// 6. Distill results into our internal segment model
	segments := make([]models.Segment, 0, len(output.Segments))
	for i, s := range output.Segments {
		segments = append(segments, models.Segment{
			ID:        fmt.Sprintf("%d", i+1),
			StartTime: s.Start,
			EndTime:   s.End,
			Text:      strings.TrimSpace(s.Text),
		})
	}

	return segments, nil
}
