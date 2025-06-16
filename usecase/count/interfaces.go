package urlCount

type Service interface {
	FetchURLHitCount(shortURL string) (int, error)
}

type Database interface {
	FetchURLHitCount(shortURL string) (int, error)
}
