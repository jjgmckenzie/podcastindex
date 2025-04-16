package episode

import "net/url"

type TranscriptType string

const (
	TranscriptPlaintext      TranscriptType = "text/plain"
	TranscriptHTML           TranscriptType = "text/html"
	TranscriptVtt            TranscriptType = "text/vtt"
	TranscriptApplicationSRT TranscriptType = "application/srt"
	TranscriptTextSRT        TranscriptType = "text/srt"
	TranscriptJSON           TranscriptType = "application/json"
)

type Transcript struct {
	URL  url.URL
	Type TranscriptType
}
