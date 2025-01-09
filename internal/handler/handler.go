package handler

import "github.com/wxlbd/nunu-casbin-admin/internal/service"

type Handler struct {
	user *UserHandler
	role *RoleHandler
	menu *MenuHandler
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		user: NewUserHandler(svc),
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
