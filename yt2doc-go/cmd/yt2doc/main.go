package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/samar-108/yt2doc-go/pkg/formatter"
	"github.com/samar-108/yt2doc-go/pkg/segmenter"
)

func main() {
	// 1. Setup Input: URL and Output File
	videoURL := flag.String("url", "", "YouTube video URL")
	outputFile := flag.String("o", "output.md", "Output file path")
	flag.Parse()

	if *videoURL == "" {
		log.Fatal("Usage: yt2doc -url <youtube-url>")
	}

	ctx := context.Background()

	// 2. Step One: Segmenting (Download + Transcribe + Chunk)
	fmt.Println(" Starting pipeline: Download and Transcription...")
	segSvc := segmenter.NewSegmenter(ctx)
	doc, err := segSvc.Segment(segmenter.SegmenterParams{URL: *videoURL})
	if err != nil {
		log.Fatalf("Segmenting failed: %v", err)
	}

	// 3. Step Two: Formatting (Final Output)
	fmt.Printf("Generating Markdown: %s\n", *outputFile)
	fmtSvc, err := formatter.NewFormatter()
	if err != nil {
		log.Fatalf("Failed to initialize Formatter: %v", err)
	}

	if err := fmtSvc.ToMarkdown(doc, *outputFile); err != nil {
		log.Fatalf("Formatting failed: %v", err)
	}

	fmt.Println(" Done! Your document is ready.")
}
