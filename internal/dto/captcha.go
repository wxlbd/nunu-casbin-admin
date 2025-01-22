package dto

type CaptchaResponse struct {
	CaptchaId    string `json:"captcha_id"`    // 验证码ID
	CaptchaImage string `json:"captcha_image"` // Base64编码的验证码图片
}
