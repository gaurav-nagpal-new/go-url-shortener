package dto

type URLShortenRequestBody struct {
	OriginalURL string `json:"original_url"`
	URLLength   int    `json:"url_length"`
}

type URLShortenResponseBody struct {
	ShortURL string `json:"short_url,omitempty"`
	Code     int    `json:"code"`
	Error    string `json:"error,omitempty"`
}
