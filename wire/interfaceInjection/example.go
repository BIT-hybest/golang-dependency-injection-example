package interfaceInjection

import "github.com/google/wire"

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar struct {
	X string
}

func provideBar(f Fooer) Bar {
	// f will be a *MyFooer.
	return Bar{X: f.Foo()}
}

var Set = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideBar)
