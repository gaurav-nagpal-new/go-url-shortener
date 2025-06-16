package shortenurl

type Service interface {
	CreateShortenURL(string, int) (string, error)
	FetchOriginalURL(string) (string, error)
}

type Database interface {
	CreateShortenURL(string, string) (string, error)
	FetchOriginalURL(string) (string, error)
}
