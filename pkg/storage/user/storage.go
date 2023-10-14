package user

type Storage interface {
	GetUserState(id int64) (string, error)
	SetUserState(id int64, state string) error
	CreateUser(id int64) error
}
