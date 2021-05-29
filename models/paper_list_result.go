package models

// PaperListResult is the result of PaperListResult() function.
type PaperListResult struct {
	Count    int64                 `json:"count"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Results  []PaperListResultItem `json:"results"`
}

type PaperListResultItem struct {
	ID               string   `json:"id"`
	ArxivID          *string  `json:"arxiv_id"`
	NipsID           *string  `json:"nips_id"`
	URLAbs           string   `json:"url_abs"`
	URLPDF           string   `json:"url_pdf"`
	Title            string   `json:"title"`
	Abstract         string   `json:"abstract"`
	Authors          []string `json:"authors"`
	Published        YyyyMmDdDashed   `json:"published"`
	Conference       *string  `json:"conference"`
	ConferenceURLAbs *string  `json:"conference_url_abs"`
	ConferenceURLPDF *string  `json:"conference_url_pdf"`
	Proceeding       *string  `json:"proceeding"`
}