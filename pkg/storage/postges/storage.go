package postges

import (
	postgresRoom "Telegram/pkg/storage/postges/room"
	postgresUser "Telegram/pkg/storage/postges/user"
	"Telegram/pkg/storage/room"
	"Telegram/pkg/storage/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
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
	db, err := gorm.Open(postgres.Open(s.cfg.DSN))
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

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
	return postgresUser.NewStorage(s.db)
}

func (s *Storage) GetRoomStorage() room.Storage {
	return postgresRoom.NewStorage(s.db)
}
