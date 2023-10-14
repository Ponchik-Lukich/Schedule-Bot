package ydb

import (
	"Telegram/pkg/storage/room"
	"Telegram/pkg/storage/user"
	ydbRoom "Telegram/pkg/storage/ydb/room"
	ydbUser "Telegram/pkg/storage/ydb/user"
	"context"
	ydb "github.com/PotatoHD404/gorm-driver"
	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"gorm.io/gorm"
)

type Storage struct {
	cfg Config
	db  *gorm.DB
}

func NewStorage(cfg Config) *Storage {
	return &Storage{cfg: cfg}
}

func (s *Storage) Init() *gorm.DB {
	return s.db
}

func (s *Storage) Connect() error {
	ctx := context.Background()

	db, err := gorm.Open(ydb.Open(s.cfg.Database,
		ydb.With(
			environ.WithEnvironCredentials(ctx),
		)),
	)

	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Storage) Close() error {
	if s.db != nil {
		sqlDB, err := s.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

func (s *Storage) GetUserStorage() user.Storage {
	return ydbUser.NewStorage(s.db)
}

func (s *Storage) GetRoomStorage() room.Storage {
	return ydbRoom.NewStorage(s.db)
}
