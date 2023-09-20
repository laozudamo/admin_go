package request

type Login struct {
	UserName  string `json:"userName" binding:"required"`  // 用户名
	Password  string `json:"password" binding:"required"`  // 密码
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

type UserCrete struct {
	UserName string  `json:"userName" binding:"required"` // 用户名
	Password string  `json:"password" binding:"required"` // 密码
	RoleKey  RoleKey `json:"roleKey" binding:"required"`  // 角色
}

type RoleKey = uint

const (
	Admin RoleKey = 0
	User  RoleKey = 1
)

type UserUpdate struct {
	ID       uint    `json:"id" binding:"required"`       // 用户ID
	Password string  `json:"password" binding:"required"` // 密码
	RoleKey  RoleKey `json:"roleKey" binding:"required"`  // 角色
}
