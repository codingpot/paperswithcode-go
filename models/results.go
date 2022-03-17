package models

// Result is the evaluation result from a paper.
type Result struct {
	ID                 string            `json:"id"`
	BestRank           *int              `json:"best_rank"`
	Metrics            map[string]string `json:"metrics"`
	Methodology        string            `json:"methodology"`
	UsesAdditionalData bool              `json:"uses_additional_data"`
	Paper              *string           `json:"paper"`
	BestMetric         *string           `json:"best_metric"`
	EvaluatedOn        YyyyMmDdDashed    `json:"evaluated_on"`
	ExternalSourceUrl  *string           `json:"external_source_url"`
}

// ResultList contains results produced in a paper.
type ResultList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Result
}
