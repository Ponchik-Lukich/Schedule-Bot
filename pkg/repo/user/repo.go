package user

import (
	"Telegram/pkg/storage/user"
)

type Repository interface {
	CreateUser(id int64) error
	GetUserState(id int64) (string, error)
	SetUserState(id int64, state string) error
}

type repository struct {
	storage user.Storage
}

func NewRepository(storage user.Storage) Repository {
	return &repository{storage: storage}
}

func (r *repository) CreateUser(id int64) error {
	return r.storage.CreateUser(id)
}

func (r *repository) GetUserState(id int64) (string, error) {
	return r.storage.GetUserState(id)
}

func (r *repository) SetUserState(id int64, state string) error {
	return r.storage.SetUserState(id, state)
}
