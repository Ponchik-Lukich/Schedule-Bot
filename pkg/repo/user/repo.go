package user

import (
	"Telegram/pkg/storage/ydb/user"
)

type Repository interface {
}

type repository struct {
	storage user.Storage
}

func NewRepository(storage user.Storage) Repository {
	return &repository{storage: storage}
}
