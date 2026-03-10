package chaptering

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/samar-108/yt2doc-go/internal/models"
	"google.golang.org/api/option"
)

type Chapterer struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

type ChapterParams struct {
	Document *models.Document
}

type aiResponse struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func NewChapterer(ctx context.Context, apiKey string) (*Chapterer, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	model := client.GenerativeModel("gemini-2.0-flash-exp")
	// Hint: We tell the model to respond in JSON format
	model.ResponseMIMEType = "application/json"

	return &Chapterer{client: client, model: model}, nil
}

func (c *Chapterer) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

func (c *Chapterer) Chapter(ctx context.Context, params ChapterParams) (*models.Document, error) {
	doc := params.Document

	for i := range doc.Chapters {
		// 1. Build the transcript for this specific chapter
		var sb strings.Builder
		for _, seg := range doc.Chapters[i].Segments {
			sb.WriteString(seg.Text + " ")
		}
		transcript := sb.String()

		// 2. Define the prompt
		prompt := fmt.Sprintf(`Summarize this transcript. 
		Return JSON with:
		- "title": A concise heading for this section.
		- "summary": A 2-sentence summary of the main points.
		
		Transcript: %s`, transcript)

		// 3. Call Gemini
		resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			return nil, fmt.Errorf("AI generation failed for chapter %d: %w", i, err)
		}

		// 4. Parse the AI response
		if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
			part := resp.Candidates[0].Content.Parts[0]
			if text, ok := part.(genai.Text); ok {
				var aiResult aiResponse
				if err := json.Unmarshal([]byte(text), &aiResult); err != nil {
					// Fallback: If JSON parsing fails, just save the raw text as summary
					doc.Chapters[i].Summary = string(text)
					continue
				}
				doc.Chapters[i].Title = aiResult.Title
				doc.Chapters[i].Summary = aiResult.Summary
			}
		}
	}

	return doc, nil
}
