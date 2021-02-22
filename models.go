package vctr

type Permissions int

const (
	PermViewLinks Permissions = 1 << (iota + 1)
	PermCreateLinks
	PermUpdateLinks
	PermDeleteLinks
	PermViewUsers
	PermCreateUsers
	PermUpdateUsers
	PermDeleteUsers
	PermPerformStateChanges
	PermCreateApiKey

	PermUnset Permissions = -1
	PermAdmin Permissions = 2147483647
)

type CountModel struct {
	Count int `json:"count"`
}

type EntityModel struct {
	Guid    string `json:"guid"`
	Created Time   `json:"created"`
}

type UserModel struct {
	*EntityModel

	UserName    string      `json:"username"`
	Permissions Permissions `json:"permissions"`
	LastLogin   Time        `json:"last_login"`
}

type UserCreateModel struct {
	UserName    string      `json:"username"`
	Password    string      `json:"password"`
	Permissions Permissions `json:"permissions"`
}

type UserLoginModel struct {
	*UserModel

	SessionKey string `json:"session_key"`
}

type UpdateSelfUserModel struct {
	UserName        string `json:"username"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type LinkModel struct {
	*EntityModel

	Ident             string `json:"ident"`
	Destination       string `json:"destination"`
	Enabled           bool   `json:"enabled"`
	PermanentRedirect bool   `json:"permanent_redirect"`
	PasswordRequired  bool   `json:"password_required"`
	LastAccess        Time   `json:"last_access"`
	AccessCount       int    `json:"access_count"`
	UniqueAccessCount int    `json:"unique_access_count"`
	TotalAccessLimit  int    `json:"total_access_limit"`
	Expires           Time   `json:"expires"`

	Creator *UserModel `json:"creator"`
}

type LinkCreateModel struct {
	Ident             string `json:"ident"`
	Destination       string `json:"destination"`
	Enabled           bool   `json:"enabled"`
	PermanentRedirect bool   `json:"permanent_redirect"`
	TotalAccessLimit  int    `json:"total_access_limit"`
	Expires           Time   `json:"expires"`
	Password          string `json:"password"`
}

type LoginModel struct {
	Ident    string `json:"ident"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type ApiKeyModel struct {
	*EntityModel

	LastAccess  Time `json:"last_access"`
	AccessCount int  `json:"access_count"`
}

type ApiKeyCreatedModel struct {
	*ApiKeyModel

	Key string `json:"key"`
}
