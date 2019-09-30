package main

import "github.com/spf13/viper"

type (
	// Handler is an interface each commnad have to implement
	// to get the input
	Handler interface {
		Handle(config *viper.Viper) error
	}
)
