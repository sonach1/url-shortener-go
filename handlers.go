package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"urlshortener/models"
	"urlshortener/utils"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)

	shortCode := utils.GenerateShortCode()
	models.SaveURL(shortCode, req.URL)
	models.CacheURL(shortCode, req.URL)

	resp := map[string]string{"short_url": "http://localhost:8080/" + shortCode}
	json.NewEncoder(w).Encode(resp)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	url, err := models.GetCachedURL(code)
	if err != nil {
		url, err = models.GetOriginalURL(code)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		models.CacheURL(code, url)
	}
	http.Redirect(w, r, url, http.StatusFound)
}
