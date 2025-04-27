package episode

import (
	"encoding/json"
	"net/url"
)

type TranscriptType string

const (
	TranscriptPlaintext      TranscriptType = "text/plain"
	TranscriptHTML           TranscriptType = "text/html"
	TranscriptVTT            TranscriptType = "text/vtt"
	TranscriptApplicationSRT TranscriptType = "application/srt"
	TranscriptTextSRT        TranscriptType = "text/srt"
	TranscriptJSON           TranscriptType = "application/json"
)

type Transcript struct {
	URL  url.URL
	Type TranscriptType
}

func (t Transcript) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		URL  string         `json:"URL"`
		Type TranscriptType `json:"Type"`
	}{
		URL:  t.URL.String(),
		Type: t.Type,
	})
}

func (t *Transcript) UnmarshalJSON(data []byte) error {
	aux := &struct {
		URL  string         `json:"URL"`
		Type TranscriptType `json:"Type"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	parsedURL, err := url.Parse(aux.URL)
	if err != nil {
		return err
	}

	t.URL = *parsedURL
	t.Type = aux.Type
	return nil
}
