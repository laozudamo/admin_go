package request

type Login struct {
	UserName  string `json:"userName" binding:"required"`  // 用户名
	Password  string `json:"password" binding:"required"`  // 密码
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}
