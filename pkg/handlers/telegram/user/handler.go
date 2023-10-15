package user

import (
	"Telegram/pkg/repo/user"
)

type Handler interface {
}

type handler struct {
	repo user.Repository
}

func NewHandler(repo user.Repository) Handler {
	return &handler{repo: repo}
}
