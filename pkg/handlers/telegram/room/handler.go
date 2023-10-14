package room

import (
	"Telegram/pkg/repo/room"
)

type Handler interface {
}

type handler struct {
	repo room.Repository
}

func NewHandler(repo room.Repository) Handler {
	return &handler{repo: repo}
}
