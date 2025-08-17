package pokeapi

import (
	"bytes"
	"encoding/json" // You'll need this import for json.NewDecoder
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetResponse(pageURL *string) ([]byte, error) {

	req, err := http.NewRequest("GET", *pageURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) DecodeIntoJson(rawdata []byte, target interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(rawdata))
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("error decoding JSON response: %w", err)
	}
	return nil
}

func (c *Client) GetJsonResponseAndDecode(pageURL *string, target interface{}) error {

	req, err := http.NewRequest("GET", *pageURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Always good to check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-OK status: %d %s", resp.StatusCode, resp.Status)
	}

	// This is the key part: decoding into the 'target' interface
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("error decoding JSON response: %w", err)
	}

	return nil
}
