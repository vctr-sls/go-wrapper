package vctr

type userClient struct {
	*Client
}

func (c *userClient) GetMe() (res *UserModel, err error) {
	res = new(UserModel)
	err = c.r.Get("users/me", nil, res)
	return
}
