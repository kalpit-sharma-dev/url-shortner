package model

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	ID      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	TinyUrl string `json:"tinyUrl,omitempty"`
	Domain  string `json:"domain,omitempty"`
}
