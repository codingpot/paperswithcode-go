package paperswithcode_go

import "encoding/json"

func (c *Client) sendGetRequest(url string, result interface{}) error {
	response, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}

	return json.NewDecoder(response.Body).Decode(result)
}
