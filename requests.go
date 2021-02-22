package vctr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type requestClient struct {
	rootEndpoint string
	authHeader   string
	client       *http.Client
}

func newRequestClient(rootEndpoint string, authHeader ...string) (c *requestClient) {
	c = new(requestClient)

	if !strings.HasPrefix(rootEndpoint, "https://") && !strings.HasPrefix(rootEndpoint, "http://") {
		rootEndpoint = "https://" + rootEndpoint
	}

	c.rootEndpoint = strings.TrimSuffix(rootEndpoint, "/")
	c.client = http.DefaultClient
	if len(authHeader) > 0 && authHeader[0] != "" {
		c.authHeader = authHeader[0]
	}
	return
}

func (c *requestClient) SetAuthHeader(header string) {
	c.authHeader = header
}

func (c *requestClient) Get(path string, query url.Values, res interface{}) error {
	return c.request("GET", path, nil, query, res)
}

func (c *requestClient) Post(path string, body interface{}, query url.Values, res interface{}) error {
	return c.request("POST", path, body, query, res)
}

func (c *requestClient) Delete(path string, query url.Values) error {
	return c.request("DELETE", path, nil, query, nil)
}

func (c *requestClient) request(
	method, path string,
	body interface{},
	query url.Values,
	response interface{},
) (err error) {
	resource, err := url.Parse(fmt.Sprintf("%s/api/%s", c.rootEndpoint, strings.TrimPrefix(path, "/")))
	if err != nil {
		return
	}
	if query != nil {
		resource.RawQuery = query.Encode()
	}

	var bodyReader io.ReadWriter
	if body != nil {
		bodyReader = bytes.NewBuffer([]byte{})
		if err = json.NewEncoder(bodyReader).Encode(body); err != nil {
			return
		}
	}

	req, err := http.NewRequest(method, resource.String(), bodyReader)
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")
	if c.authHeader != "" {
		req.Header.Add("authorization", c.authHeader)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode >= 400 {
		err = parseError(res)
		return
	}

	if response != nil {
		err = json.NewDecoder(res.Body).Decode(response)
	}

	return
}

func parseError(res *http.Response) (err *ResponseError) {
	err = new(ResponseError)
	err.Code = res.StatusCode
	err.Message = res.Status

	resBody := new(struct {
		Error  string `json:"error"`
		Errors map[string][]string
	})
	if parseErr := json.NewDecoder(res.Body).Decode(resBody); parseErr != nil {
		return
	}

	if resBody.Error != "" {
		err.Message = resBody.Error
	} else if resBody.Errors != nil && len(resBody.Errors) > 0 {
		errors := make([]string, len(resBody.Errors))
		i := 0
		for k, v := range resBody.Errors {
			errors[i] = k + ": " + v[0]
			i++
		}
		err.Message = strings.Join(errors, "\n")
	}

	return
}
