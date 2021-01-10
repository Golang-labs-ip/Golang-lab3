//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Golang-labs-ip/Golang-lab3/tablets"
)

func ComposeApiServer(port HttpPortNumber) (*TabletApiServer, error) {
	wire.Build(
		NewDbConnection,
		tablets.Providers,
		wire.Struct(new(TabletApiServer), "Port", "TabletsHandler"),
	)
	return nil, nil
}
