package repos

import (
	"encoding/json"
	"log"
	"net/http"
	"stonks/internal/models"
)

type NewsRepo struct {
	client *http.Client
}

func (r *NewsRepo) GetNews(req *http.Request) (models.News, error) {
	resp, err := r.client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Print(json.NewDecoder(resp.Body))
		return models.News{}, err
	}

	newsBody := models.News{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newsBody)
	if err != nil {
		log.Print(err.Error())
	}
	return newsBody, nil
}
