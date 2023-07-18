package guards

type User struct {
	id         int
	permission bool
}

func NewUser(id int, permission bool) *User {
	return &User{id: id, permission: permission}
}

type Guard struct {
	user *User
}

func (guard *Guard) HasPermissions() bool {
	return guard.user.permission
}

func NewGuard(user *User) *Guard {
	return &Guard{user: user}
}
