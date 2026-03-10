package downloader

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/samar-108/yt2doc-go/internal/models"
	"go.uber.org/zap"
)

// Download pulls the audio and metadata from a given URL.
// It returns the path to the downloaded audio file and the document metadata.
// NOTE: The caller is responsible for deleting the returned audio file when finished.
func Download(ctx context.Context, url string) (string, *models.Document, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Create a temporary directory for the download
	tmpDir, err := os.MkdirTemp("", "yt2doc-*")
	if err != nil {
		return "", nil, fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Define output file path
	// We use .m4a as it's a common high-quality audio-only format from YouTube
	audioPath := filepath.Join(tmpDir, "audio.m4a")

	// 1. Run yt-dlp to get video info (JSON)
	logger.Info("Fetching video info", zap.String("url", url))
	infoCmd := exec.CommandContext(ctx, "yt-dlp", "-j", "--skip-download", url)
	infoOut, err := infoCmd.Output()
	if err != nil {
		os.RemoveAll(tmpDir) // Cleanup on error
		return "", nil, fmt.Errorf("failed to fetch video info: %w", err)
	}

	// 2. Parse video info
	var dlpResp models.DlpResponse
	if err := json.Unmarshal(infoOut, &dlpResp); err != nil {
		os.RemoveAll(tmpDir)
		return "", nil, fmt.Errorf("failed to parse video info: %w", err)
	}

	// 3. Create document metadata
	publishedAt, err := time.Parse("20060102", dlpResp.UploadDate)
	if err != nil {
		publishedAt = time.Now()
	}

	doc := &models.Document{
		MetaData: models.MetaData{
			Title:       dlpResp.Title,
			Description: dlpResp.Description,
			Author:      dlpResp.Uploader,
			PublishedAt: publishedAt,
			Duration:    dlpResp.Duration,
			Thumbnail:   dlpResp.Thumbnail,
			VideoID:     dlpResp.ID,
			Url:         dlpResp.WebpageURL,
		},
	}

	// 4. Download best audio
	logger.Info("Downloading audio", zap.String("url", url))

	// -x: Extract audio
	// --audio-format m4a: Specify format
	// -o: Output template
	downloadCmd := exec.CommandContext(ctx, "yt-dlp",
		"-x",
		"--audio-format", "m4a",
		"-o", audioPath,
		url,
	)

	// Capture stderr for better debugging
	var stderr strings.Builder
	downloadCmd.Stderr = &stderr

	if err := downloadCmd.Run(); err != nil {
		os.RemoveAll(tmpDir)
		return "", nil, fmt.Errorf("failed to download audio: %v (stderr: %s)", err, stderr.String())
	}

	return audioPath, doc, nil
}
