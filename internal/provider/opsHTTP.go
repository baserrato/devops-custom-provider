package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetOp(op_id string) (*Ops_Api, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/op/%s", c.HostURL, op_id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(res)
	if err != nil {
		return nil, err
	}
	op := Ops_Api{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &op)
	if err != nil {
		return nil, err
	}

	return &op, nil
}

func (c *Client) CreateOp(dev Ops_Api) (*Ops_Api, error) {
	rb, err := json.Marshal(dev)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/op", c.HostURL), strings.NewReader(string(rb)))
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
func (c *Client) UpdateOp(op Ops_Api) (*Ops_Api, error) {
	rb, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/op/%s", c.HostURL, op.Id), strings.NewReader(string(rb)))
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

func (c *Client) DeleteOp(Id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/op/%s", c.HostURL, Id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
