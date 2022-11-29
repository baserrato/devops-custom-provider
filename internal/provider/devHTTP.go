package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/baserrato/devops-resource"
)

func (c *Client) GetDev(dev_id string) (*devops_resource.Dev, error) {
	//req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers", c.HostURL), nil)
	res, err := http.NewRequest("GET", fmt.Sprintf("%s/dev/id/%s", c.HostURL, dev_id), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(res)
	if err != nil {
		return nil, err
	}
	dev := devops_resource.Dev{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &dev)
	if err != nil {
		return nil, err
	}

	return &dev, nil
}

func (c *Client) CreateDev(dev devops_resource.Dev) (*devops_resource.Dev, error) {
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

	newDev := devops_resource.Dev{}
	err = json.Unmarshal(body, &newDev)
	if err != nil {
		return nil, err
	}

	return &newDev, nil
}
func (c *Client) UpdateDev(dev devops_resource.Dev) (*devops_resource.Dev, error) {
	rb, err := json.Marshal(dev)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/dev/%s", c.HostURL, dev.Id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newDev := devops_resource.Dev{}
	err = json.Unmarshal(body, &newDev)
	if err != nil {
		return nil, err
	}

	return &newDev, nil
}

func (c *Client) DeleteDev(Id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/dev/%s", c.HostURL, Id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
