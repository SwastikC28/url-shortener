package model

type URLData struct {
	ShortURL         string   `json:"short_url"`
	LongURL          string   `json:"long_url"`
	AccessCount      int      `json:"access_count"`
	AccessTimestamps []string `json:"access_timestamps"`
}
