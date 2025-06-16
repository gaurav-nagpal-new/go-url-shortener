package web

import (
	urlCount "go-url-shortener/usecase/count"
	"net/http"
)

type hitCountAPI struct {
	service urlCount.Service
}

func NewHitCountAPI(service urlCount.Service) *hitCountAPI {
	return &hitCountAPI{
		service: service,
	}
}

func (h *hitCountAPI) FetchURLHitCount(w http.ResponseWriter, r *http.Request) {
	// fetch from mongoDB
}
