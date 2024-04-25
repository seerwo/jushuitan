package erp

import (
	"github.com/seerwo/jushuitan/credential"
	"github.com/seerwo/jushuitan/erp/config"
	"github.com/seerwo/jushuitan/erp/context"
	"github.com/seerwo/jushuitan/erp/oauth"
	"github.com/seerwo/jushuitan/erp/order"
)

// erp relate api
type Erp struct {
	ctx *context.Context
}

//NewOpenPlatform new openplatform
func NewErp(cfg *config.Config) *Erp {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, cfg.AuthCode, credential.CacheKeyOfficialAccountPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Erp{ctx: ctx}
}

//SetAccessTokenHandle custom access_token get method
func (y *Erp) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	y.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (y *Erp) GetContext() *context.Context {
	return y.ctx
}

//GetAccessToken get access_token
func (y *Erp) GetAccessToken() (string, error) {
	return y.ctx.GetAccessToken()
}

// GetOauth oauth2 web oauth
func (y *Erp) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(y.ctx)
}

// GetOrder get order
func (y *Erp) GetOrder() *order.Order {
	return order.NewOrder(y.ctx)
}
