package eval

import (
	"github.com/killthebuddh4/gecko/types"
)

var Array types.Exec = func(scope *types.Trajectory, arguments ...types.Value) (types.Value, error) {
	return arguments, nil
}
