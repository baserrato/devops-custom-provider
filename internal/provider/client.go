package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:8080"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
}

// NewClient -
func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{},
		// Default  URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 201 && res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
