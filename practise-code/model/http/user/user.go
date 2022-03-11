package user

// User register structure
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User login structure
type LoginRequest struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

// User info structure
type UserInfoResponse struct {
	ID     uint  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
	HeaderImg string `json:"headerImg" `
}

