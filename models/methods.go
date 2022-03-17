package models

type Method struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	FullName    string  `json:"full_name"`
	Description string  `json:"description"`
	Paper       *string `json:"paper"`
}

// MethodList represents methods used in the paper.
type MethodList struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Method `json:"results"`
}
