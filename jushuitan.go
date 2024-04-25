package jushuitan

import (
	"github.com/seerwo/jushuitan/cache"
	"github.com/seerwo/jushuitan/erp"
	"github.com/seerwo/jushuitan/erp/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Jushuitan struct
type Jushuitan struct {
	cache cache.Cache
}

// NewJushuitan init
func NewJushuitan() *Jushuitan {
	return &Jushuitan{}
}

//SetCache  set cache
func (ym *Jushuitan) SetCache(cahce cache.Cache) {
	ym.cache = cahce
}

//GetOfficialAccount get erp
func (ym *Jushuitan) GetErp(cfg *config.Config) *erp.Erp {
	if cfg.Cache == nil {
		cfg.Cache = ym.cache
	}
	return erp.NewErp(cfg)
}