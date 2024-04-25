package context

import (
	"github.com/seerwo/jushuitan/credential"
	"github.com/seerwo/jushuitan/erp/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}

