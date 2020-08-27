//+build wireinject

package structInjection

import "github.com/google/wire"

func injectFooBar() *FooBar {
	wire.Build(Set)
	return &FooBar{}
}
