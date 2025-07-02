package userwebrequest

type UserUpdatePassword struct {
	Id       int
	Password string `json:"-"`
}
