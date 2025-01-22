package handler

import (
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
)

type Handler struct {
	user    *UserHandler
	role    *RoleHandler
	menu    *MenuHandler
	dict    *DictHandler
	captcha *CaptchaHandler
	cfg     *config.Config
}

func NewHandler(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		user:    NewUserHandler(svc, cfg),
		role:    NewRoleHandler(svc),
		menu:    NewMenuHandler(svc),
		dict:    NewDictHandler(svc.Dict()),
		captcha: NewCaptchaHandler(svc),
		cfg:     cfg,
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

func (h *Handler) Captcha() *CaptchaHandler {
	return h.captcha
}
