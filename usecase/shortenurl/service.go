package shortenurl

import (
	"fmt"
	"math/rand"
	"os"
)

type service struct {
	sqliteDB Database
}

func NewService(db Database) Service {
	return &service{
		sqliteDB: db,
	}
}

func (s *service) CreateShortenURL(originalURL string, URLLength int) (string, error) {
	// creat a random string of 6 characters from the alphabets

	subURL := ""

	for range URLLength {
		subURL += string(rand.Intn(26) + 97)
	}

	shortenurl := fmt.Sprintf("%s%s", os.Getenv("HOST"), subURL)
	return s.sqliteDB.CreateShortenURL(originalURL, shortenurl)
}

func (s *service) FetchOriginalURL(shortURL string) (string, error) {
	return s.sqliteDB.FetchOriginalURL(shortURL)
}
