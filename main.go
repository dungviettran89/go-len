package main

import (
	"fmt"
	. "github.com/dungviettran89/go-len/context"
	"github.com/gogap/aop"
)

type Config struct {
	Prefix string
}

func (config *Config) Before(jp aop.JoinPoint) {
	fmt.Printf("Before...")
}

func (config *Config) After(jp aop.JoinPoint) {
	fmt.Printf("After...")
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
