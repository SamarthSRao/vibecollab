package models

import (
	"time"
)

type Segment struct {
	ID        string  `json:"id"`
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
	Text      string  `json:"text"`
}

type Chapter struct {
	Title     string    `json:"title"`
	Segments  []Segment `json:"segments"`
	StartTime float64   `json:"start_time"`
	EndTime   float64   `json:"end_time"`
	Summary   string    `json:"summary"`
}

type MetaData struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Duration    float64   `json:"duration"`
	Thumbnail   string    `json:"thumbnail"`
	VideoID     string    `json:"video_id"`
	Url         string    `json:"url"`
}

type Document struct {
	MetaData MetaData  `json:"metadata"`
	Chapters []Chapter `json:"chapters"`
}

type DlpResponse struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Uploader     string   `json:"uploader"`
	UploaderID   string   `json:"uploader_id"`
	ViewCount    int64    `json:"view_count"`
	LikeCount    int64    `json:"like_count"`
	DislikeCount int64    `json:"dislike_count"`
	Duration     float64  `json:"duration"`
	UploadDate   string   `json:"upload_date"`
	Thumbnail    string   `json:"thumbnail"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Categories   []string `json:"categories"`
	WebpageURL   string   `json:"webpage_url"`
	Formats      []Format `json:"formats"`
}

type Format struct {
	FormatID   string  `json:"format_id"`
	FormatNote string  `json:"format_note"`
	Ext        string  `json:"ext"`
	Vcodec     string  `json:"vcodec"`
	Acodec     string  `json:"acodec"`
	Resolution string  `json:"resolution"`
	FPS        float64 `json:"fps"`
	ABR        float64 `json:"abr"`
	VBR        float64 `json:"vbr"`
	Tbr        float64 `json:"tbr"`
	Tfmt       string  `json:"tfmt"`
	Url        string  `json:"url"`
}
