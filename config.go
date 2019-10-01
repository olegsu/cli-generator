package main

import "github.com/spf13/viper"

type (
	// Handler is an interface each command have to implement
	// to get the input
	Config interface{}
)

func NewConfig() Config {
	return viper.New()
}
