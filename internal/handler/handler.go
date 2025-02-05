package handler

import (
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
)

type Handler struct {
	user    *UserHandler
	role    *RoleHandler
	dict    *DictHandler
	captcha *CaptchaHandler
	sysMenu *SysMenuHandler
	cfg     *config.Config
}

func NewHandler(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		user:    NewUserHandler(svc, cfg),
		role:    NewRoleHandler(svc),
		dict:    NewDictHandler(svc.Dict()),
		captcha: NewCaptchaHandler(svc),
		sysMenu: NewSysMenuHandler(svc),
		cfg:     cfg,
	}
}

func (h *Handler) User() *UserHandler {
	return h.user
}

func (h *Handler) Role() *RoleHandler {
	return h.role
}

func (h *Handler) Dict() *DictHandler {
	return h.dict
}

func (h *Handler) Captcha() *CaptchaHandler {
	return h.captcha
}

func (h *Handler) SysMenu() *SysMenuHandler {
	return h.sysMenu
}
