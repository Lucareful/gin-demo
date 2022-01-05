package request

// CreateUserRequest  user 表 create 接口请求 struct
type CreateUserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required,max=100"`
}

// UpdateUserRequest  user 表 update 接口请求 struct.
type UpdateUserRequest struct {
	Id       uint   `json:"id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required,max=100"`
}
