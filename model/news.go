package model

type News struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Content string `json:"description"`
	Source  string `json:"source"`
	Country string `json:"country"`
}
