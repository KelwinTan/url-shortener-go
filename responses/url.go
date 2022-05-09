package responses

import "github.com/KelwinTan/url-shortener-go/models"

type ShortenURLResponse struct {
	Url        models.Url `json:"url"`
	ShortenURL string     `json:"shorten_url"`
}
