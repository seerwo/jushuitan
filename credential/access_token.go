package credential

//AccessTokenHandle AccessToken interface
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
	GetAccessParam(method string, req interface{})(accessParam string, err error)
}

