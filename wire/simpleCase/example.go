package simpleCase

import (
	"context"
	"errors"
	"github.com/google/wire"
)

type Foo struct {
	X int
}

// ProvideFoo constructor of Foo
func ProvideFoo() Foo {
	return Foo{X: 42}
}

type Bar struct {
	X int
}

// ProvideBar constructor of Bar
func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

type Baz struct {
	X int
}

// ProvideBaz constructor of Bar
func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

var SuperSet = wire.NewSet(ProvideBar, ProvideFoo, ProvideBaz)

func CreateBaz(ctx context.Context) (Baz, error) {
	foo := ProvideFoo()
	bar := ProvideBar(foo)
	baz, err := ProvideBaz(ctx, bar)
	if err != nil {
		return Baz{}, err
	}
	return baz, nil
}

