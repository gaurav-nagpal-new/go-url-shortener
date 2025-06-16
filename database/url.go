package database

import (
	"go-url-shortener/model"
	"os"
)

func (u *SQLClient) CreateShortenURL(originalURL, shortURL string) (string, error) {

	url := &model.URL{
		OriginalURL:  originalURL,
		ShortenedURL: shortURL,
	}

	res := u.Client.Create(url)
	if res.Error != nil {
		return "", res.Error
	}

	return url.ShortenedURL, nil
}

func (u *SQLClient) FetchOriginalURL(shortURL string) (string, error) {

	var url *model.URL
	shortURL = os.Getenv("HOST") + shortURL
	res := u.Client.First(&url, "shortened_url = ?", shortURL)

	if res.Error != nil {
		return "", res.Error
	}

	return url.ShortenedURL, nil
}
