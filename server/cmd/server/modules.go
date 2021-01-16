//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Golang-labs-ip/Golang-lab3/server/tablets"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeAPIServer(port HTTPPortNumber) (*TabletAPIServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from balancers package.
		tablets.Providers,
		// Provide BalancerApiServer instantiating the structure and injecting balancers handler and port number.
		wire.Struct(new(TabletAPIServer), "Port", "tabletsHandler"),
	)
	return nil, nil
}
