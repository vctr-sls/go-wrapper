package vctr

import "net/url"

type Client struct {
	r *requestClient

	Users *userClient
}

func NewClient(endpoint string, authHeader ...string) (c *Client) {
	c = new(Client)
	c.r = newRequestClient(endpoint, authHeader...)

	c.Users = &userClient{c}

	return
}

func NewClientWithApiToken(endpoint, token string) (c *Client) {
	c = NewClient(endpoint)
	c.SetApiToken(token)
	return
}

func NewClientWithSessionToken(endpoint, token string) (c *Client) {
	c = NewClient(endpoint)
	c.SetSessionToken(token)
	return
}

func (c *Client) SetSessionToken(token string) {
	c.r.SetAuthHeader("session " + token)
}

func (c *Client) SetApiToken(token string) {
	c.r.SetAuthHeader("basic " + token)
}

func (c *Client) Login(login *LoginModel) (res *UserLoginModel, err error) {
	res = new(UserLoginModel)
	err = c.r.Post("auth/login", login, url.Values{
		"getSessionKey": []string{"true"},
	}, res)
	if err != nil {
		return
	}

	c.SetSessionToken(res.SessionKey)

	return
}
