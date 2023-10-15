package repo

import (
	"Telegram/pkg/repo/room"
	"Telegram/pkg/repo/user"
	"Telegram/pkg/storage"
)

type Repositories interface {
	GetUserRepo() user.Repository
	GetRoomRepo() room.Repository
}

type repositories struct {
	storage storage.Storage
}

func NewRepositories(storage storage.Storage) Repositories {
	return &repositories{
		storage: storage,
	}
}

func (r *repositories) GetUserRepo() user.Repository {
	return user.NewRepository(r.storage.GetUserStorage())
}

func (r *repositories) GetRoomRepo() room.Repository {
	return room.NewRepository(r.storage.GetRoomStorage())
}
