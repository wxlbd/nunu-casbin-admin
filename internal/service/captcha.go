package service

import (
	"context"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/wxlbd/gin-casbin-admin/internal/handler"
)

type captchaService struct {
	store base64Captcha.Store
}

func NewCaptchaService(redisClient *redis.Client) handler.CaptchaService {
	return &captchaService{
		store: NewRedisStore(redisClient, 10*time.Minute),
	}
}

func (s *captchaService) Generate(ctx context.Context) (id, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, s.store)
	id, b64s, _, err = captcha.Generate()
	return
}

func (s *captchaService) Verify(ctx context.Context, id, answer string) bool {
	return s.store.Verify(id, answer, true)
}
