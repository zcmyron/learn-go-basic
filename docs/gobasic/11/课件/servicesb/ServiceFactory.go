package servicesb

import "com.jtthink/core"

type ServiceFactory struct {

}

func NewServiceFactory()  *ServiceFactory{
	return  &ServiceFactory{}
}

func (sf *ServiceFactory) Create(name string)  core.IService {

	switch name {
	case "news":
		return &NewsService{}
	case "user":
		return &UserService{}
	default:
		return nil

	}
}
