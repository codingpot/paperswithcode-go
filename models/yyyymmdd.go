package models

import (
	"encoding/json"
	"strings"
	"time"
)

type YyyyMmDdDashed time.Time

func (y *YyyyMmDdDashed) UnmarshalJSON(bytes []byte) error {
	cleanString := strings.Trim(string(bytes), `"`)
	parsedTime, err := time.Parse("2006-01-02", cleanString)
	if err != nil {
		return err
	}
	*y = YyyyMmDdDashed(parsedTime)
	return nil
}

func (y YyyyMmDdDashed) MarshalJSON() ([]byte, error) {
	return json.Marshal(y)
}

func NewYyyyMmDdDashed(year int, month time.Month, day int) YyyyMmDdDashed {
	return YyyyMmDdDashed(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}
