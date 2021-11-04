package request

// ListArticleRequest article 表 list 接口请求 struct.
type ListArticleRequest struct {
	Name  string `form:"name" binding:"required"`
	State int    `form:"state" binding:"eq=0"`
}

// CreateArticleRequest article 表 create 接口请求 struct.
type CreateArticleRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatedBy string `json:"created_by" binding:"required,max=100"`
	State     int    `json:"state" binding:"eq=0|eq=1"`
}

// UpdateArticleRequest article 表 update 接口请求 struct.
type UpdateArticleRequest struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}
