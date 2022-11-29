package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/liatrio/devops-bootcamp/examples/ch6/devops-resources"
)

func (c *Client) GetDevOps() ([]devops_resource.DevOps, error) {
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
	devops := []devops_resource.DevOps{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &devops)
	if err != nil {
		return nil, err
	}

	return devops, nil
}
func (c *Client) GetDevOp(devop_id string) (*devops_resource.DevOps, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/devops/%s", c.HostURL, devop_id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(res)
	if err != nil {
		return nil, err
	}
	devops := devops_resource.DevOps{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &devops)
	if err != nil {
		return nil, err
	}

	return &devops, nil
}

func (c *Client) CreateDevOps(devops devops_resource.DevOps) (*devops_resource.DevOps, error) {
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

	newDevOp := devops_resource.DevOps{}
	err = json.Unmarshal(body, &newDevOp)
	if err != nil {
		return nil, err
	}

	return &newDevOp, nil
}

func (c *Client) UpdateDevOps(devops devops_resource.DevOps) (*devops_resource.DevOps, error) {
	rb, err := json.Marshal(devops)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/devops/%s", c.HostURL, devops.Id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newDevOps := devops_resource.DevOps{}
	err = json.Unmarshal(body, &newDevOps)
	if err != nil {
		return nil, err
	}

	return &newDevOps, nil
}

func (c *Client) DeleteDevOps(Id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/devops/%s", c.HostURL, Id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
