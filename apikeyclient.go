package vctr

type apikeyClient struct {
	*Client
}

// Generate creates or re-creates a new API key
// and returns the key details as well as the
// key string itself.
//
// Attention: API keys are only returned on
// (re-)generation and never again after that.
func (c *apikeyClient) Generate() (res *ApiKeyCreatedModel, err error) {
	res = new(ApiKeyCreatedModel)
	err = c.r.Post("apikey", nil, nil, res)
	return
}

// Details returns the properties of the currently
// generated API key (or 404 error it not exist)
// without the actual key string.
func (c *apikeyClient) Details() (res *ApiKeyModel, err error) {
	res = new(ApiKeyModel)
	err = c.r.Get("apikey", nil, res)
	return
}

// Delete resets the generated API key and makes
// it unsuable for subsequent authentication tries.
//
// The delete action returns no response, so it is
// considered to be successful, when err != nil.
func (c *apikeyClient) Delete() (err error) {
	err = c.r.Delete("apikey", nil)
	return
}
