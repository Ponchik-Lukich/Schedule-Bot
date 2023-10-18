package storage

import (
	cst "Telegram/pkg/constants"
	errorsMes "Telegram/pkg/errors"
	"Telegram/pkg/storage/postges"
	"Telegram/pkg/storage/room"
	"Telegram/pkg/storage/user"
	"errors"
	"gorm.io/gorm"
)

type Storage interface {
	Connect() error
	Close() error
	Init() *gorm.DB

	GetUserStorage() user.Storage
	GetRoomStorage() room.Storage
}

func NewStorage(dbType string, cfg Config) (Storage, error) {
	switch dbType {
	case cst.Postgres:
		if postgresCfg, ok := cfg.(*postges.Config); ok {
			return postges.NewStorage(*postgresCfg), nil
		} else {
			return nil, errors.New(errorsMes.InvalidConfig)
		}
	default:
		return nil, errors.New(errorsMes.UnsupportedDatabaseType)
	}
}
