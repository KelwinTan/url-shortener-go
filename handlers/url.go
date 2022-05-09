package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/KelwinTan/url-shortener-go/config"
	"github.com/KelwinTan/url-shortener-go/models"
	"github.com/KelwinTan/url-shortener-go/requests"
	"github.com/KelwinTan/url-shortener-go/responses"
	"github.com/teris-io/shortid"
)

func (h DBHandler) Default(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, []string{"URL Shortener"})
}

func (h DBHandler) GetUrls(w http.ResponseWriter, r *http.Request) {
	var urls []models.Url

	if result := h.DB.Find(&urls); result.Error != nil {
		log.Error(result.Error)
		responses.ERROR(w, http.StatusUnprocessableEntity, result.Error)
	} else {
		responses.JSON(w, http.StatusOK, urls)
	}
}

func (h DBHandler) RedirectUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var url models.Url

	if result := h.DB.First(&url, "short_url = ?", vars["short_url"]); result.Error != nil {
		log.Error(result.Error)
		responses.ERROR(w, http.StatusNotFound, result.Error)
	} else {
		http.Redirect(w, r, url.Url, http.StatusFound)
	}
}

func (h DBHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var shortenUrlRequest requests.ShortenUrlRequest
	json.Unmarshal(body, &shortenUrlRequest)

	// to check if the url is valid
	_, err = url.ParseRequestURI(shortenUrlRequest.Url)
	if err != nil {
		log.Error(err)
	}

	shortUrl, err := shortid.Generate()
	if err != nil {
		log.Error(err)
	}

	url := models.Url{
		Url:      shortenUrlRequest.Url,
		ShortUrl: shortUrl,
	}

	if result := h.DB.Create(&url); result.Error != nil {
		log.Error(result.Error)
		responses.ERROR(w, http.StatusUnprocessableEntity, result.Error)
	} else {
		responses.JSON(w, http.StatusCreated, responses.ShortenURLResponse{
			Url:        url,
			ShortenURL: fmt.Sprintf("localhost:%s/%s", config.GetConfig().AppPort, url.ShortUrl),
		})
	}
}

func (h DBHandler) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var url models.Url

	if result := h.DB.First(&url, "short_url = ?", vars["short_url"]); result.Error != nil {
		log.Error(result.Error)
		responses.ERROR(w, http.StatusNotFound, result.Error)
	} else {
		h.DB.Delete(&url)
		responses.JSON(w, http.StatusOK, url)
	}
}

// user can update their own custom short url
func (h DBHandler) UpdateUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var updatedUrl requests.UpdateShortUrlRequest
	json.Unmarshal(body, &updatedUrl)

	var url models.Url
	if result := h.DB.First(&url, "short_url = ?", vars["short_url"]); result.Error != nil {
		log.Error(result.Error)
		responses.ERROR(w, http.StatusNotFound, result.Error)
	} else {
		url.Url = updatedUrl.Url
		url.ShortUrl = updatedUrl.CustomShortUrl

		h.DB.Save(url)
		responses.JSON(w, http.StatusOK, url)
	}
}
