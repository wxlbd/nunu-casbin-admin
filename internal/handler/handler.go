package handler

import (
	"github.com/wxlbd/nunu-casbin-admin/internal/service"
	"github.com/wxlbd/nunu-casbin-admin/pkg/config"
)

type Handler struct {
	user *UserHandler
	role *RoleHandler
	menu *MenuHandler
	cfg  *config.Config
}

func NewHandler(svc service.Service, cfg *config.Config) *Handler {
	return &Handler{
		cfg:  cfg,
		user: NewUserHandler(svc, cfg),
		role: NewRoleHandler(svc),
		menu: NewMenuHandler(svc),
	}
}

func (h *Handler) User() *UserHandler {
	return h.user
}

func (h *Handler) Role() *RoleHandler {
	return h.role
}

func (h *Handler) Menu() *MenuHandler {
	return h.menu
}
