package SibylSystem

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (s *SibylSystem) Get(method string, params url.Values) (json.RawMessage, error) {
	r, err := http.NewRequest("GET", s.ApiUrl+method, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build GET request to %s: %w", method, err)
	}
	r.URL.RawQuery = params.Encode()
	var client http.Client
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request to %s: %w", method, err)
	}
	defer resp.Body.Close()

	var b Body

	if err = json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return nil, fmt.Errorf("failed to decode GET request to %s: %w", method, err)
	}
	if !b.Success {
		return nil, &b.Error
	}
	return b.Result, nil
}
