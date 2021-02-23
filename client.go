package vctr

// Client provides a simple interface to access
// the vctr REST API gateway.
type Client struct {
	// Authorization endpoints.
	Auth *authClient
	// Users endpoints.
	Users *userClient
	// Links endpoints.
	Links *linksClient
	// API Key endpoints.
	ApiKey *apikeyClient

	r *requestClient
}

// NewClient creates a new client instance connecting
// to the passed vctr instance endpoint, for example
// https://s.zekro.de or https://vctr.yourhost.com.
//
// You can also pass optional authorization header tokens
// which can be used to authenticate against the API, but
// it is recommented to use the methods NewClientWithApiToken
// and NewClientWithSessionToken for that purpose.
func NewClient(endpoint string, authHeader ...string) (c *Client) {
	c = new(Client)
	c.r = newRequestClient(endpoint, authHeader...)

	c.Auth = &authClient{c}
	c.Users = &userClient{c}
	c.Links = &linksClient{c}
	c.ApiKey = &apikeyClient{c}

	return
}

// NewClientWithApiKey is shorthand for NewClient with
// an initial API key used to authenticate against the
// API.
func NewClientWithApiToken(endpoint, token string) (c *Client) {
	c = NewClient(endpoint)
	c.SetApiToken(token)
	return
}

// NewClientWithSessionKey is shorthand for NewClient with
// an initial session key used to authenticate against the
// API.
func NewClientWithSessionToken(endpoint, token string) (c *Client) {
	c = NewClient(endpoint)
	c.SetSessionToken(token)
	return
}

// SetApiToken sets the passed API token to the
// client so that it will be used in subsequent requests
// to authenticate against the API.
func (c *Client) SetApiToken(token string) {
	c.r.SetAuthHeader("basic " + token)
}

// SetSessionToken sets the passed session token to the
// client so that it will be used in subsequent requests
// to authenticate against the API.
func (c *Client) SetSessionToken(token string) {
	c.r.SetAuthHeader("session " + token)
}
