package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/baserrato/devops-resource"
)

func (c *Client) GetEngineers() ([]devops_resource.Engineer, error) {
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
	engineers := []devops_resource.Engineer{}
	//var results map[string]interface{}
	err = json.Unmarshal(body, &engineers)
	if err != nil {
		return nil, err
	}

	return engineers, nil
}

func (c *Client) CreateEngineer(engineer devops_resource.Engineer) (*devops_resource.Engineer, error) {
	rb, err := json.Marshal(engineer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/engineers", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	newEngineer := devops_resource.Engineer{}
	err = json.Unmarshal(body, &newEngineer)
	if err != nil {
		return nil, err
	}

	return &newEngineer, nil
}

func (c *Client) UpdateEngineer(engineer devops_resource.Engineer) (*devops_resource.Engineer, error) {
	rb, err := json.Marshal(engineer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/engineers/%s", c.HostURL, engineer.Id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	newEngineer := devops_resource.Engineer{}
	err = json.Unmarshal(body, &newEngineer)
	if err != nil {
		return nil, err
	}

	return &newEngineer, nil
}

func (c *Client) DeleteEngineer(Id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/engineers/%s", c.HostURL, Id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetEngineer(Id string) (*devops_resource.Engineer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers/id/%s", c.HostURL, Id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	newEngineer := devops_resource.Engineer{}
	err = json.Unmarshal(body, &newEngineer)
	if err != nil {
		return nil, err
	}

	return &newEngineer, nil
}

func (c *Client) GetEngineerWithName(Name string) (*devops_resource.Engineer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/engineers/name/%s", c.HostURL, Name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	newEngineer := devops_resource.Engineer{}
	err = json.Unmarshal(body, &newEngineer)
	if err != nil {
		return nil, err
	}

	return &newEngineer, nil
}
