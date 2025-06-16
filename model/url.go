package model

type URL struct {
	ID           int    `gorm:"primary_key"`
	OriginalURL  string `gorm:"not null"` // long URL
	ShortenedURL string `gorm:"not null"` // short URL after applying logic
}
