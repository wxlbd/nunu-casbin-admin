package handler

import (
	"github.com/wxlbd/gin-casbin-admin/internal/service"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
)

type Handler struct {
	user *UserHandler
	role *RoleHandler
	menu *MenuHandler
	dict *DictHandler
	cfg  *config.Config
}

func NewHandler(svc service.Service, cfg *config.Config) *Handler {
	return &Handler{
		cfg:  cfg,
		user: NewUserHandler(svc, cfg),
		role: NewRoleHandler(svc),
		menu: NewMenuHandler(svc),
		dict: NewDictHandler(svc.Dict()),
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

func (h *Handler) Dict() *DictHandler {
	return h.dict
}
