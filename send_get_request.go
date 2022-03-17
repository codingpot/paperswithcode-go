package paperswithcode_go

import "encoding/json"

func getJson[T any](client *Client, url string) (T, error) {
	var result T
	response, err := client.httpClient.Get(url)

	if err != nil {
		return result, err
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
