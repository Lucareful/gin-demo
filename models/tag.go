package models

// Tag tag 表的结构体
type Tag struct {
	Model

	Name       string `json:"name" validate:"required"`
	CreatedBy  string `json:"created_by" validate:"required,max=100"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state" validate:"eq=0|eq=1"`
}

// GetTags 获取指定页数的 tag 表记录
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// GetTagTotal 获取 tag 表所有记录条数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName 判断 tag 记录是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

// AddTag 为 tag 表增加一条记录
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}
