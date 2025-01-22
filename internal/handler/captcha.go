package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type CaptchaHandler struct {
	svc Service
}

func NewCaptchaHandler(svc Service) *CaptchaHandler {
	return &CaptchaHandler{
		svc: svc,
	}
}

// Generate 生成验证码
// @Summary 生成验证码
// @Description 生成图片验证码
// @Tags 认证管理
// @Accept json
// @Produce json
// @Success 200 {object} ginx.Response{data=dto.CaptchaResponse} "成功"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Router /auth/captcha [get]
func (h *CaptchaHandler) Generate(c *gin.Context) {
	id, b64s, err := h.svc.Captcha().Generate(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, &dto.CaptchaResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	})
}
