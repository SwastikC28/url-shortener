package command

type CreateShortenURLCommand struct {
	LongURL     string `json:"long_url"`
	CustomAlias string `json:"custom_alias"`
	TTL         int    `json:"ttl_seconds"`
}

type UpdateShortURLCommand struct{}
