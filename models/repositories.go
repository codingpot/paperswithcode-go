package models

// RepositoryList contains code repositories implemented by a certain paper.
type RepositoryList struct {
	Count    int64        `json:"count"`
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []Repository `json:"results"`
}

type Repository struct {
	URL         string `json:"url"`
	IsOfficial  bool   `json:"is_official"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Framework   string `json:"framework"`
}
