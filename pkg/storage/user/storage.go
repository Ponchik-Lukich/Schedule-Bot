package user

type Storage interface {
	GetUserState(id int64) (string, error)
	CreateUser(id int64) error
	UpdateUser(id int64, updates map[string]interface{}) error
}
