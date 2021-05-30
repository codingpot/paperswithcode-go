package models

import "fmt"

// PaperListResult is the result of PaperListResult() function.
type PaperListResult struct {
	Count    int64                 `json:"count"`
	Next     *string               `json:"next"`
	Previous *string               `json:"previous"`
	Results  []PaperListResultItem `json:"results"`
}

var _ fmt.Stringer = (*PaperListResultItem)(nil)

type PaperListResultItem struct {
	ID               string         `json:"id"`
	ArxivID          *string        `json:"arxiv_id"`
	NipsID           *string        `json:"nips_id"`
	URLAbs           string         `json:"url_abs"`
	URLPDF           string         `json:"url_pdf"`
	Title            string         `json:"title"`
	Abstract         string         `json:"abstract"`
	Authors          []string       `json:"authors"`
	Published        YyyyMmDdDashed `json:"published"`
	Conference       *string        `json:"conference"`
	ConferenceURLAbs *string        `json:"conference_url_abs"`
	ConferenceURLPDF *string        `json:"conference_url_pdf"`
	Proceeding       *string        `json:"proceeding"`
}

func (p PaperListResultItem) String() string {
	return fmt.Sprintf("ID: %s ArxivID: %v NipsID: %v, URLAbs: %s, URLPDF: %s Title: %s Authors: %s",
		p.ID, p.ArxivID, p.NipsID, p.URLAbs, p.URLPDF, p.Title, p.Authors)
}