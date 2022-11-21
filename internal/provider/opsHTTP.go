package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetOp(Id string) (*Ops_Api, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/ops/%s", c.HostURL, Id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(res)
	if err != nil {
		return nil, err
	}

	ops := Ops_Api{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &ops)
	if err != nil {
		return nil, err
	}

	return &ops, nil
}

func (c *Client) UpdateOps(ops Ops_Api) (*Ops_Api, error) {
	rb, err := json.Marshal(ops)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ops/", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newOps := Ops_Api{}
	err = json.Unmarshal(body, &newOps)
	if err != nil {
		return nil, err
	}

	return &newOps, nil
}

func (c *Client) CreateOp(ops Ops_Api) (*Ops_Api, error) {
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

	newOp := Ops_Api{}
	err = json.Unmarshal(body, &newOp)
	if err != nil {
		return nil, err
	}

	return &newOp, nil
}

func (c *Client) DeleteOps(Id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/ops/%s", c.HostURL, Id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
