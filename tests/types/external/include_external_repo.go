//go:generate go run github.com/antoninferrand/pergolator github.com/antoninferrand/pergolator/tests/types/external.IncludeExternalRepo

package external

import (
	"time"

	"github.com/dave/jennifer/jen"
)

type IncludeExternalRepo struct {
	Time  time.Time `json:"time"`
	Group jen.Group `json:"group"`
}
