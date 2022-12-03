package model

type Request struct {
	ID      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Domain  string `json:"domain,omitempty"`
	TinyUrl string `json:"tinyUrl,omitempty"`
	//CreatedAt time.Time `json:"createdAt,omitempty"`
}

type Response struct {
	ID      int    `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	TinyUrl string `json:"tinyUrl,omitempty"`
	Domain  string `json:"domain,omitempty"`
	//CreatedAt time.Time `json:"createdAt,omitempty"`
}
