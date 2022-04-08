package endpoints

type GetRequest struct {
	Shortname string `json:"shortname,omitempty"`
}

type GetResponse struct {
	OriginalURL string `json:"original_url"`
	Err         string `json:"err,omitempty"`
}

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	Shortname string `json:"shortname"`
	Err       string `json:"err,omitempty"`
}
