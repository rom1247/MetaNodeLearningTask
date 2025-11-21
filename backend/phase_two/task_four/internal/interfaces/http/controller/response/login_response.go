package response

type LoginResponse struct {
	// 用户ID
	UserID uint `json:"userId"`
	// 用户名
	Username string `json:"username"`
	// Token
	Token string `json:"token"`
	// Token 过期时间 单位秒
	Expire int64 `json:"expire" `
}
