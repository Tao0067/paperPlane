package request

// LoginRequest 登陆请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//type loginParams struct {
//	Name     string `json:"name" binding:"required"`
//	Password string `json:"password" binding:"required"`
//}
