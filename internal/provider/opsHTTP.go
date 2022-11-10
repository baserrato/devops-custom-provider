package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) GetOps() ([]Ops, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.Get("http://localhost:8080/ops")
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
	ops := []Ops{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &ops)
	if err != nil {
		return nil, err
	}

	return ops, nil
}

func (c *Client) CreateOp(ops Ops) (*Ops, error) {
	rb, err := json.Marshal(ops)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ops", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newOp := Ops{}
	err = json.Unmarshal(body, &newOp)
	if err != nil {
		return nil, err
	}

	return &newOp, nil
}
