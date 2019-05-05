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
	appContext.RegisterBean(
		NewConfig,
		Advice{MethodName: "Before", Execution: "PrintPrefix()", Ordering: aop.Before},
		Advice{MethodName: "After", Execution: "PrintPrefix()", Ordering: aop.After},
	)
	err := appContext.Invoke(func(config *Config) {
		config.PrintPrefix()
	})
	if err != nil {
		panic(err)
	}
}
