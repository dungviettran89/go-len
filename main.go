package main

import (
	"fmt"
	. "go-len/context"
)

type Config struct {
	Prefix string
}

func (config *Config) PrintPrefix() {
	fmt.Printf(config.Prefix)
}

func NewConfig() *Config {
	return &Config{
		Prefix: "test",
	}
}

func main() {
	appContext := New()
	err := appContext.RegisterBean(
		NewConfig,
	)
	if err != nil {
		panic(err)
	}
	err = appContext.InvokeFunc(func(config *Config) {
		config.PrintPrefix()
	})
	if err != nil {
		panic(err)
	}
}
