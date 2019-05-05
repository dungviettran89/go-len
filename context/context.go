package context

import (
	"go.uber.org/dig"
)

type ApplicationContext struct {
	*dig.Container
}

func New() *ApplicationContext {
	digContainer := dig.New()
	appContext := &ApplicationContext{
		digContainer,
	}

	return appContext
}

func (appContext *ApplicationContext) RegisterBean(constructor interface{}) error {
	err := appContext.Provide(constructor)
	return err
}

func (appContext *ApplicationContext) InvokeFunc(function interface{}) error {
	err := appContext.Invoke(function)
	return err
}
