package request

// ListTagRequest tag 表 list 接口请求 struct
type ListTagRequest struct {
	Name  string `json:"name" validate:"required"`
	State int    `json:"state" validate:"eq=0|eq=1"`
}

// GetTagRequest tag 表 get 接口请求 struct
type GetTagRequest struct {
	Name       string `json:"name" validate:"required"`
	CreatedBy  string `json:"created_by" validate:"required,max=100"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state" validate:"eq=0|eq=1"`
}

// CreateTagRequest tag 表 create 接口请求 struct
type CreateTagRequest struct {
}

// UpdateTagRequest tag 表 update 接口请求 struct
type UpdateTagRequest struct {
}

// DeleteTagRequest tag 表 update 接口请求 struct
type DeleteTagRequest struct {
}
