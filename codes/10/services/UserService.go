package services

type UserService struct {
}

func (ns *UserService) Get(id int) string {
	return "user1"
}
