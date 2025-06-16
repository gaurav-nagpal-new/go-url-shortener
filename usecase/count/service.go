package urlCount

type service struct {
	Db Database
}

func NewService(db Database) Service {
	return &service{
		Db: db,
	}
}

func (s *service) FetchURLHitCount(shortenURL string) (int, error) {
	return 0, nil
}
