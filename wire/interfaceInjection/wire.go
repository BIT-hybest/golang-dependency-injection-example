//+build wireinject

package interfaceInjection

import "github.com/google/wire"

func initializeBaz() (Bar, error) {
	wire.Build(Set)
	return Bar{}, nil
}
