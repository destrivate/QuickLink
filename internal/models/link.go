package models

type Link struct {
	OriginalPath string `json:"OriginalPath"`
	Path         string `json:"Path"`
	Password     string `json:"Password"`
	Redirected   int
}
