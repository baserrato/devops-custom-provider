package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) GetDevOps() ([]DevOps_Api, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.Get("http://localhost:8080/devops")
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
	devops := []DevOps_Api{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &devops)
	if err != nil {
		return nil, err
	}

	return devops, nil
}

func (c *Client) CreateDevOp(devops DevOps_Api) (*DevOps_Api, error) {
	rb, err := json.Marshal(devops)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/devops", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newDevOp := DevOps_Api{}
	err = json.Unmarshal(body, &newDevOp)
	if err != nil {
		return nil, err
	}

	return &newDevOp, nil
}
