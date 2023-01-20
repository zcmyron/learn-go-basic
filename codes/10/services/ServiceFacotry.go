package services

type ServiceFactory struct {
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{}
}

func (sf *ServiceFactory) Create(name string) IService {
	switch name {
	case "news":
		return &NewsService{}
	case "user":
		return &UserService{}
	default:
		return nil
	}
}
