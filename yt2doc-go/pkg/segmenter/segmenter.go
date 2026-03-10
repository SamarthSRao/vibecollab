package segmenter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/samar-108/yt2doc-go/internal/models"
	"github.com/samar-108/yt2doc-go/pkg/downloader"
	"github.com/samar-108/yt2doc-go/pkg/transcriber"
)

type Segmenter struct {
	ctx context.Context
}
type SegmenterParams struct {
	URL string
}

func NewSegmenter(ctx context.Context) *Segmenter {
	return &Segmenter{ctx: ctx}
}

// Segment executes the full pipeline: Download -> Transcribe -> Group into Chapters
func (s *Segmenter) Segment(params SegmenterParams) (*models.Document, error) {
	// 1. Download
	audioPath, doc, err := downloader.Download(s.ctx, params.URL)
	if err != nil {
		return nil, fmt.Errorf("download failed: %w", err)
	}
	// Cleanup audio folder after processing
	defer os.RemoveAll(filepath.Dir(audioPath))

	// 2. Transcribe
	segments, err := transcriber.Transcribe(s.ctx, audioPath)
	if err != nil {
		return nil, fmt.Errorf("transcription failed: %w", err)
	}

	// 3. Grouping Logic (The actual "Segmenting")
	// We'll group segments into 60-second chunks to create "Chapters"
	doc.Chapters = s.groupSegmentsIntoChapters(segments, 60.0)

	return doc, nil
}

// groupSegmentsIntoChapters takes raw segments and chunks them by a maximum duration (in seconds)
func (s *Segmenter) groupSegmentsIntoChapters(segments []models.Segment, maxDuration float64) []models.Chapter {
	if len(segments) == 0 {
		return nil
	}

	var chapters []models.Chapter
	var currentChapter models.Chapter

	chapterStartTime := segments[0].StartTime
	currentChapter.StartTime = chapterStartTime

	for i, seg := range segments {
		currentChapter.Segments = append(currentChapter.Segments, seg)
		currentChapter.EndTime = seg.EndTime

		// If this segment pushes us past the duration, or if it's the last segment, close the chapter
		duration := seg.EndTime - chapterStartTime
		if duration >= maxDuration || i == len(segments)-1 {
			currentChapter.Title = fmt.Sprintf("Section starting at %.2f", chapterStartTime)
			chapters = append(chapters, currentChapter)

			// Start a new chapter
			if i < len(segments)-1 {
				currentChapter = models.Chapter{}
				chapterStartTime = segments[i+1].StartTime
				currentChapter.StartTime = chapterStartTime
			}
		}
	}

	return chapters
}
