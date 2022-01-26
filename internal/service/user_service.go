package service

type IUserInterface interface {
	AddUser(int, string) error
}

type User struct {
	IUser IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "user created")
}
