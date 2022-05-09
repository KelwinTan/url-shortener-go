package requests

type ShortenUrlRequest struct {
	Url string `json:"url"`
}

type UpdateShortUrlRequest struct {
	Url            string `json:"url"`
	CustomShortUrl string `json:"custom_short_url"`
}
