package models

type UserModel struct {
}

func (user *UserModel) Get(id int) string {
	return "User1"
}
