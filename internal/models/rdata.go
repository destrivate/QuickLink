package models

type RequestData struct {
	Path string `json:"Path"`
}

type RequestDelete struct {
	Path     string `json:"Path"`
	Password string `json:"Password"`
}
