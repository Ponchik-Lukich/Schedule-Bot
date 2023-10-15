package storage

import (
	"Telegram/pkg/constants"
	errorsMes "Telegram/pkg/errors"
	"Telegram/pkg/storage/room"
	"Telegram/pkg/storage/user"
	"Telegram/pkg/storage/ydb"
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
	case constants.Ydb:
		if ydbCfg, ok := cfg.(*ydb.Config); ok {
			return ydb.NewStorage(*ydbCfg), nil
		} else {
			return nil, errors.New(errorsMes.InvalidConfig)
		}
	default:
		return nil, errors.New(errorsMes.UnsupportedDatabaseType)
	}
}
