package servicesb

type UserService struct {
}

func (us *UserService) Get(id int) string {
	return "user1b"
}
