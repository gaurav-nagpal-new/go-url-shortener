package web

import (
	"encoding/json"
	"go-url-shortener/dto"
	"go-url-shortener/usecase/shortenurl"
	"go-url-shortener/utils"
	"net/http"
)

type shortnURLAPI struct {
	service shortenurl.Service
}

func NewShortenURLAPI(service shortenurl.Service) shortnURLAPI {
	return shortnURLAPI{
		service: service,
	}
}

func (a *shortnURLAPI) CreateShortenURLHander(w http.ResponseWriter, r *http.Request) {
	// decode the request body
	var requestBody *dto.URLShortenRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		// send 400 bad request
		utils.SendResponse(&dto.URLShortenResponseBody{
			Code:  http.StatusBadRequest,
			Error: "invalid request body",
		}, http.StatusBadRequest, w)
		return
	}

	// validation if url is not present
	if requestBody.OriginalURL == "" {
		utils.SendResponse(&dto.URLShortenResponseBody{
			Code:  http.StatusBadRequest,
			Error: "original_url/url_length cannot be empty",
		}, http.StatusBadRequest, w)
		return
	}

	// call the method from service to shorten the URL, and send the response
	shortened, err := a.service.CreateShortenURL(requestBody.OriginalURL, requestBody.URLLength)
	if err != nil {
		// send 500 Response
		utils.SendResponse(&dto.URLShortenResponseBody{
			Code:  http.StatusInternalServerError,
			Error: "unable to process the request",
		}, http.StatusInternalServerError, w)
		return
	}

	// send the shortenedURL in response
	utils.SendResponse(&dto.URLShortenResponseBody{
		Code:     http.StatusOK,
		ShortURL: shortened,
	}, http.StatusOK, w)

}

func (a *shortnURLAPI) FetchOriginalURLHandler(w http.ResponseWriter, r *http.Request) {

	// accept shortURL in query params
	rQ := r.URL.Query()
	shortenurl := rQ.Get("shorten_url")
	originalURL, err := a.service.FetchOriginalURL(shortenurl)
	if err != nil {
		// send 500 Response
		utils.SendResponse(&dto.URLShortenResponseBody{
			Code:  http.StatusInternalServerError,
			Error: "unable to process the request",
		}, http.StatusInternalServerError, w)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)

	// increase hit count in mongoDB
	
}
