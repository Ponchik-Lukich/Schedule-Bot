package storage

import (
	"Telegram/pkg/constant"
	errorsMes "Telegram/pkg/errors"
	"Telegram/pkg/storage/ydb"
	"errors"
	"gorm.io/gorm"
)

type Storage interface {
	Connect() error
	Close() error
	Init() *gorm.DB
}

func NewStorage(dbType string, cfg Config) (Storage, error) {
	switch dbType {
	case constant.Ydb:
		if ydbCfg, ok := cfg.(*ydb.Config); ok {
			return ydb.NewStorage(*ydbCfg), nil
		} else {
			return nil, errors.New(errorsMes.InvalidConfig)
		}
	default:
		return nil, errors.New(errorsMes.UnsupportedDatabaseType)
	}
}
