package models

import (
	"time"

	"github.com/luenci/go-gin-example/types/request"

	"github.com/jinzhu/gorm"
)

// Tag tag 表的结构体
type Tag struct {
	gorm.Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
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
func ExistTagByName(name string) error {
	var tag Tag
	return db.Select("id").Where("name = ?", name).First(&tag).Error
}

// AddTag 为 tag 表增加一条记录
func AddTag(r request.CreateTagRequest) error {
	return db.Create(&Tag{
		Name:      r.Name,
		State:     r.State,
		CreatedBy: r.CreatedBy,
	}).Error
}

// BeforeCreate 创建 tag 记录时候设置 CreatedOn 字段为随着时间自增
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedOn", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}

// BeforeUpdate 更新 tag 记录时候设置 CreatedOn 字段为随着时间更新
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("ModifiedOn", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
