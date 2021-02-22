package vctr

type userClient struct {
	*Client
}

// List returns a list of users.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *userClient) List(limit, offset int) (res []*UserModel, err error) {
	res = make([]*UserModel, 0)
	err = c.r.Get("users", getLimitOffsetQuery(limit, offset), &res)
	for _, r := range res {
		r.hydrate(c.Client)
	}
	return
}

// Create user creates a new user with the given
// users properties and returns the created user
// entity.
func (c *userClient) Create(user *UserCreateModel) (res *UserModel, err error) {
	res = new(UserModel)
	err = c.r.Post("users", user, nil, res)
	res.hydrate(c.Client)
	return
}

// Me returns the user information of the
// currently authenticated user.
func (c *userClient) Me() (res *UserModel, err error) {
	return c.Get("me")
}

// UpdateMe updates the currently authenticated user
// by the given newMe properties and returns the
// new me user entity.
func (c *userClient) UpdateMe(newMe *UserUpdateSelfModel) (res *UserModel, err error) {
	res = new(UserModel)
	err = c.r.Post("users/me", newMe, nil, res)
	res.hydrate(c.Client)
	return
}

// Search searches the list of users by the
// given search query string and returns the list
// of result user entities.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *userClient) Search(query string, limit, offset int) (res []*UserModel, err error) {
	res = make([]*UserModel, 0)
	urlQuery := getLimitOffsetQuery(limit, offset)
	urlQuery.Add("query", query)
	err = c.r.Get("users/search", urlQuery, &res)
	for _, r := range res {
		r.hydrate(c.Client)
	}
	return
}

// Get returns a user entity with the requested
// user id.
func (c *userClient) Get(id string) (res *UserModel, err error) {
	res = new(UserModel)
	err = c.r.Get("users/"+id, nil, res)
	res.hydrate(c.Client)
	return
}

// Update updates the user passed by user ID
// by the given newUser properties and returns the
// new user entity.
func (c *userClient) Update(id string, newUser *UserUpdateModel) (res *UserModel, err error) {
	res = new(UserModel)
	err = c.r.Post("users/"+id, newUser, nil, res)
	res.hydrate(c.Client)
	return
}

// Delete removes the passed user by its user ID
// and all links connected to this user.
//
// The delete action returns no response, so it is
// considered to be successful, when err != nil.
func (c *userClient) Delete(id string) (err error) {
	err = c.r.Delete("users/"+id, nil)
	return
}

// MyLinks returns a list of links owned by the currently
// authenticated user.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *userClient) MyLinks(limit, offset int) (res []*LinkModel, err error) {
	return c.LinksOfUser("me", limit, offset)
}

// LinksOfUser returns a list of links owned by the given
// user by id.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *userClient) LinksOfUser(id string, limit, offset int) (res []*LinkModel, err error) {
	res = make([]*LinkModel, 0)
	err = c.r.Get("users/"+id+"/links", getLimitOffsetQuery(limit, offset), &res)
	for _, r := range res {
		r.hydrate(c.Client)
	}
	return
}

// SearchLinksMine searches through the list of links
// owned by the authenticated user and matching the given
// search query.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *userClient) SearchLinksOfMine(query string, limit, offset int) (res []*LinkModel, err error) {
	return c.SearchLinksOfUser("me", query, limit, offset)
}

// SearchLinksOfUser searches through the list of links
// owned by the given user by id and matching the given
// search query.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *userClient) SearchLinksOfUser(id, query string, limit, offset int) (res []*LinkModel, err error) {
	res = make([]*LinkModel, 0)
	urlQuery := getLimitOffsetQuery(limit, offset)
	urlQuery.Add("query", query)
	err = c.r.Get("users/"+id+"/links/search", urlQuery, &res)
	for _, r := range res {
		r.hydrate(c.Client)
	}
	return
}

// LinksOfMineCount returns the total count of links owned
// by the currently authenticated user.
func (c *userClient) LinksOfMineCount() (res *CountModel, err error) {
	return c.LinksOfUserCount("me")
}

// LinksOfUserCount returns the total count of links owned
// by the given user by id.
func (c *userClient) LinksOfUserCount(id string) (res *CountModel, err error) {
	res = new(CountModel)
	err = c.r.Get("users/"+id+"/links/count", nil, res)
	return
}

// --- MODEL FUNCTIONS ---

// Links returns a list io links of the current user entity.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (u *UserModel) Links(limit, offset int) (res []*LinkModel, err error) {
	return u.client.Users.LinksOfUser(u.Guid, limit, offset)
}

// LinksCount returns the number of links owned by
// the current user entity.
func (u *UserModel) LinksCount() (res *CountModel, err error) {
	return u.client.Users.LinksOfUserCount(u.Guid)
}

// SearchLinks searches the links of the current user
// by the given query.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (u *UserModel) SearchLinks(query string, limit, offset int) (res []*LinkModel, err error) {
	return u.client.Users.SearchLinksOfUser(u.Guid, query, limit, offset)
}
