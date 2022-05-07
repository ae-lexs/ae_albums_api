package entity

type CreateAlbumRequest struct {
	Artist string  `json:"artist"`
	Price  float64 `json:"price,string"`
	Title  string  `json:"title"`
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
