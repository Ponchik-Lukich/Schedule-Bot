package user

import (
	"Telegram/pkg/storage/user"
)

type Repository interface {
	CreateUser(id int64) error
	UpdateUser(id int64, updates map[string]any) error
	GetUserState(id int64) (string, error)
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

func (r *repository) UpdateUser(id int64, updates map[string]any) error {
	return r.storage.UpdateUser(id, updates)
}
