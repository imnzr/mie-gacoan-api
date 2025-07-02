package userwebrequest

type UserCreateRequest struct {
	Id       int
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
