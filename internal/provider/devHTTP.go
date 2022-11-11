package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) GetDevs() ([]Dev_Api, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.Get("http://localhost:8080/dev")
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
	devs := []Dev_Api{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &devs)
	if err != nil {
		return nil, err
	}

	return devs, nil
}

func (c *Client) CreateDev(dev Dev_Api) (*Dev_Api, error) {
	rb, err := json.Marshal(dev)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dev", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newDev := Dev_Api{}
	err = json.Unmarshal(body, &newDev)
	if err != nil {
		return nil, err
	}

	return &newDev, nil
}
