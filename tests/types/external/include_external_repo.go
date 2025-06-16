//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/types/external.IncludeExternalRepo

package external

import (
	"time"

	"github.com/dave/jennifer/jen"
)

type IncludeExternalRepo struct {
	Time  time.Time `json:"time"`
	Group jen.Group `json:"group"`
}
