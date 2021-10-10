package response

// ListTagResponse tag 表 list 接口响应(返回) struct
type ListTagResponse struct {
	Data map[string]interface{} `json:"data"`
}

// GetTagResponse tag 表 get 接口响应(返回) struct
type GetTagResponse struct {
}

// CreateTagResponse tag 表 create 接口响应(返回) struct
type CreateTagResponse struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	State     int    `json:"state"`
}

// UpdateTagResponse tag 表 update 接口响应(返回) struct
type UpdateTagResponse struct {
}

// DeleteTagResponse tag 表 delete 接口响应(返回) struct
type DeleteTagResponse struct {
}
