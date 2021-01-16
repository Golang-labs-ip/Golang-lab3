// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/Golang-labs-ip/Golang-lab3/server/tablets"
)

// Injectors from modules.go:

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeAPIServer(port HTTPPortNumber) (*TabletAPIServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := tablets.NewStore(db)
	httpHandlerFunc := tablets.HTTPHandler(store)
	tabletAPIServer := &TabletAPIServer{
		Port:				port,
		TabletsHandler: 	httpHandlerFunc,
	}
	return tabletAPIServer, nil
}
