package vctr

import "net/url"

type authClient struct {
	*Client
}

// When passing valid authorization credentials for an
// existing user, the authenticated user information
// is returned as well as a session key, which is set
// as session key to the current client to authenticate
// subsequent requests with it.
func (c *authClient) Login(login *LoginModel) (res *UserLoginModel, err error) {
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
