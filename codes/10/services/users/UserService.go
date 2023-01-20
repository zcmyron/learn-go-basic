package users

import core "github.com/zcmyron/learn-go-basic/codes/10/core"

type UserService struct {
}

func (us *UserService) Get(id int) string {
	return "user1"
}

func init() {
	core.SetService(NewServiceFactory().Create("user"))
}
