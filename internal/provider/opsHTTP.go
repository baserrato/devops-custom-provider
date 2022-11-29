package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	devops_resource "github.com/liatrio/devops-bootcamp/examples/ch6/devops-resources"
)

func (c *Client) GetOp(op_id string) (*devops_resource.Ops, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/op/id/%s", c.HostURL, op_id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(res)
	if err != nil {
		return nil, err
	}
	op := devops_resource.Ops{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &op)
	if err != nil {
		return nil, err
	}

	return &op, nil
}

func (c *Client) GetOpByName(op_name string) (*devops_resource.Ops, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/op/name/%s", c.HostURL, op_name), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(res)
	if err != nil {
		return nil, err
	}
	op := devops_resource.Ops{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &op)
	if err != nil {
		return nil, err
	}

	return &op, nil
}

func (c *Client) CreateOp(dev devops_resource.Ops) (*devops_resource.Ops, error) {
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

	newOp := devops_resource.Ops{}
	err = json.Unmarshal(body, &newOp)
	if err != nil {
		return nil, err
	}

	return &newOp, nil
}
func (c *Client) UpdateOp(op devops_resource.Ops) (*devops_resource.Ops, error) {
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

	newOp := devops_resource.Ops{}
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
