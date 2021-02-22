package vctr

type linksClient struct {
	*Client
}

// List returns a list of links.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *linksClient) List(limit, offset int) (res []*LinkModel, err error) {
	res = make([]*LinkModel, 0)
	err = c.r.Get("links", getLimitOffsetQuery(limit, offset), &res)
	for _, r := range res {
		r.hydrate(c.Client)
	}
	return
}

// Create creates a new link with the given properties
// and returns the resulting link entity.
func (c *linksClient) Create(link *LinkCreateModel) (res *LinkModel, err error) {
	res = new(LinkModel)
	err = c.r.Post("links", link, nil, res)
	res.hydrate(c.Client)
	return
}

// Search searches all links by the given query and
// returns a list of results.
//
// limit defines the maximum ammount of response
// entities and offset defines the number of entities
// to be skipped.
func (c *linksClient) Search(query string, limit, offset int) (res []*LinkModel, err error) {
	res = make([]*LinkModel, 0)
	urlQuery := getLimitOffsetQuery(limit, offset)
	urlQuery.Add("query", query)
	err = c.r.Get("links/search", urlQuery, &res)
	for _, r := range res {
		r.hydrate(c.Client)
	}
	return
}

// Get returns a link entity by its id.
func (c *linksClient) Get(id string) (res *LinkModel, err error) {
	res = new(LinkModel)
	err = c.r.Get("links/"+id, nil, res)
	return
}

// Update sets the given properties to the given
// link by id and returns the updated link entity model.
func (c *linksClient) Update(id string, newLink *LinkCreateModel) (res *LinkModel, err error) {
	res = new(LinkModel)
	err = c.r.Post("links/"+id, newLink, nil, res)
	return
}

// Delete removes the passed link by its
// id.
//
// The delete action returns no response, so it is
// considered to be successful, when err != nil.
func (c *linksClient) Delete(id string) (err error) {
	err = c.r.Delete("links/"+id, nil)
	return
}
