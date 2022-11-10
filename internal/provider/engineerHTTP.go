package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) GetEngineers() ([]Engineer_Api, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.Get("http://localhost:8080/engineers")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	/*
		body, err := c.doRequest(res, nil)
		if err != nil {
			return "", err
		}
	*/
	engineers := []Engineer_Api{}
	err = json.Unmarshal(body, &engineers)
	if err != nil {
		return nil, err
	}
	return engineers, nil
}

// GetCoffee - Returns specific coffee (no auth required)
func (c *Client) GetEngineer(engineerID string) ([]Engineer_Api, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers/%s", c.HostURL, engineerID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err

	}

	engineers := []Engineer_Api{}
	err = json.Unmarshal(body, &engineers)
	if err != nil {
		return nil, err
	}

	return engineers, nil
}

// CreateCoffee - Create new coffee
func (c *Client) CreateEngineer(engineer Engineer_Api, authToken *string) (*Engineer_Api, error) {
	rb, err := json.Marshal(engineer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/engineers", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	newEngineer := Engineer_Api{}
	err = json.Unmarshal(body, &newEngineer)
	if err != nil {
		return nil, err
	}

	return &newEngineer, nil
}
