package context

import (
	"github.com/gogap/aop"
	"go.uber.org/dig"
	"reflect"
)

type ApplicationContext struct {
	*aop.AOP
	*dig.Container
	aop.BeanFactory
}

func New() *ApplicationContext {
	digContainer := dig.New()
	beanFactory := aop.NewClassicBeanFactory()
	gogapAOP := aop.NewAOP()
	gogapAOP.SetBeanFactory(beanFactory)
	appContext := &ApplicationContext{
		gogapAOP,
		digContainer,
		beanFactory,
	}

	return appContext
}

type Advice struct {
	MethodName string
	Execution  string
	Ordering   aop.AdviceOrdering
}

func (appContext *ApplicationContext) RegisterBean(constructor interface{}, advices ...Advice) {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	typeOfConstructor := reflect.TypeOf(constructor)
	for i := 0; i < typeOfConstructor.NumOut(); i++ {
		outParam := typeOfConstructor.Out(i)
		beanId := outParam.String()
		appContext.BeanFactory.RegisterBean(beanId, reflect.PtrTo(outParam))
		aspect := aop.NewAspect(beanId+"_Aspect", beanId)
		aspect.SetBeanFactory(appContext.BeanFactory)
		appContext.AOP.AddAspect(aspect)
		for _, advice := range advices {
			if _, ok := typeOfConstructor.MethodByName(advice.MethodName); ok {
				var pointcutId string
				switch advice.Ordering {
				case aop.Before:
					pointcutId = beanId + "_BeforePointcut"
				case aop.After:
					pointcutId = beanId + "_AfterPointcut"
				case aop.Around:
					pointcutId = beanId + "_AroundPointcut"
				}
				pointcut := aop.NewPointcut(pointcutId).Execution(advice.Execution)
				aspect.AddPointcut(pointcut)
				aspect.AddAdvice(&aop.Advice{Ordering: advice.Ordering, Method: advice.MethodName, PointcutRefID: pointcutId})
			}
		}
	}
	err = appContext.Container.Provide(constructor)
}

func (appContext *ApplicationContext) Invoke(function interface{}) error {
	typeOfFunction := reflect.TypeOf(function)
	for i := 0; i < typeOfFunction.NumIn(); i++ {
		inParam := typeOfFunction.In(i)
		beanId := inParam.String()
		proxy, err := appContext.AOP.GetProxy(beanId)
		if err != nil {
			return err
		}
		functionValue := reflect.ValueOf(function).String()
		if name, ok := reflect.PtrTo(inParam).MethodByName(functionValue); ok {
			method := proxy.Invoke(name)
			err = appContext.Container.Invoke(method)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
