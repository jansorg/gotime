package sevdesk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: "https://my.sevdesk.de/api/v1",
		apiKey:  apiKey,
		http:    &http.Client{},
		Logging: false,
	}
}

type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
	Logging bool

	// cached data
	user              *User
	userContactPerson IDWithType
	units             *[]Unit
}

func (api *Client) LoadBasicData() error {
	if u, err := api.GetCurrentUser(); err != nil {
		return err
	} else {
		api.user = u
		api.userContactPerson = IDWithType{ID: u.ID, ObjectName: u.ObjectName}
	}

	if units, err := api.GetUnits(); err != nil {
		return err
	} else {
		api.units = &units
	}

	return nil
}

func (api *Client) doRequest(method string, path string, body io.Reader, query map[string]string) (*http.Response, error) {
	req, err := api.newRequest(method, path, body, query)
	if err != nil {
		return nil, err
	}
	return api.do(req)
}

func (api *Client) newRequest(method string, path string, body io.Reader, query map[string]string) (*http.Request, error) {
	var u *url.URL
	var err error
	if u, err = api.makeURL(path); err != nil {
		return nil, err
	}

	q := u.Query()
	if method == "GET" {
		q.Add("token", api.apiKey)
	}
	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	urlString := u.String()
	var req *http.Request
	if req, err = http.NewRequest(method, urlString, body); err != nil {
		return nil, err
	}

	if method != "GET" {
		req.Header.Add("Authorization", api.apiKey)
	}
	return req, err
}

func (api *Client) newFormRequest(method string, path string, query map[string]string, body map[string]string) (*http.Request, error) {
	req, err := api.newRequest(method, path, strings.NewReader(api.createFormValues(body, true).Encode()), query)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req, err
}

func (api *Client) doFormRequest(method string, path string, query map[string]string, body map[string]string) (*http.Response, error) {
	req, err := api.newFormRequest(method, path, query, body)
	if err != nil {
		return nil, err
	}
	return api.do(req)
}

func (api *Client) createFormValues(data map[string]string, skipEmpty bool) url.Values {
	u := url.Values{}
	for k, v := range data {
		if !skipEmpty || v != "" {
			u.Add(k, v)
		}
	}
	return u
}

func (api *Client) do(req *http.Request) (*http.Response, error) {
	if api.Logging {
		log.Printf("%s %s", req.Method, req.URL.String())
	}
	return api.http.Do(req)
}

func (api *Client) makeURL(path string) (*url.URL, error) {
	u := fmt.Sprintf("%s%s", api.baseURL, path)
	return url.Parse(u)
}

func (api *Client) unwrapJSONResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if api.Logging {
		fmt.Println(string(bytes))
	}

	// target is a pointer and will be updated
	return json.Unmarshal(bytes, &struct {
		Objects interface{} `json:"objects"`
	}{Objects: target})
}

func (api *Client) getJSON(path string, target interface{}) error {
	resp, err := api.doFormRequest("GET", path, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s : unexpected status %s", path, resp.Status)
	}
	return api.unwrapJSONResponse(resp, target)
}

func (api *Client) postForm(path string, successStates []int, postData map[string]string, response interface{}) error {
	resp, err := api.doFormRequest("POST", path, nil, postData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	found := false
	for _, s := range successStates {
		if s == resp.StatusCode {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("POST %s. Unexpected status %s", path, resp.Status)
	}
	return api.unwrapJSONResponse(resp, response)
}

func (api *Client) delete(elementType string, id string) error {
	// fixme escape ID in path?
	fullPath := fmt.Sprintf("/%s/%s", elementType, id)
	resp, err := api.doRequest("DELETE", fullPath, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("DELETE %s: unexpected status %s", fullPath, resp.Status)
	}
	return nil
}
