package request

// ListTagRequest tag 表 list 接口请求 struct
type ListTagRequest struct {
	Name  string `form:"name" binding:"required"`
	State int    `form:"state" binding:"eq=0"`
}

// CreateTagRequest tag 表 create 接口请求 struct
type CreateTagRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatedBy string `json:"created_by" binding:"required,max=100"`
	State     int    `json:"state" binding:"eq=0|eq=1"`
}

// UpdateTagRequest tag 表 update 接口请求 struct
type UpdateTagRequest struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}
