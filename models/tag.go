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

// GetTags 获取指定页数的 tag 表记录.
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// GetTag 获取单个 tag 记录.
func (t *Tag) GetTag(id uint) error {
	return db.Where("id = ?", id).Find(&Tag{}).Error
}

// UpdateTag 更新 tag 记录.
func (t *Tag) UpdateTag(r request.UpdateTagRequest) error {
	return db.Where("id = ?", r.ID).Update(&r).Error
}

// DeleteTag 删除 tag 记录.
func (t *Tag) DeleteTag(id uint) error {
	return db.Delete(&Tag{}, id).Error
}

// GetTagTotal 获取 tag 表所有记录条数.
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// ExistTagByName 判断 tag 记录是否存在.
func (t Tag) ExistTagByName(name string) error {
	return db.Select("id").Where("name = ?", name).First(&Tag{}).Error
}

// AddTag 为 tag 表增加一条记录.
func AddTag(r request.CreateTagRequest) error {
	return db.Create(&Tag{
		Name:      r.Name,
		State:     r.State,
		CreatedBy: r.CreatedBy,
	}).Error
}

// BeforeCreate 创建 tag 记录时候设置 CreatedOn 字段为随着时间自增
func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("CreatedOn", time.Now().Unix()); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate 更新 tag 记录时候设置 CreatedOn 字段为随着时间更新
func (t *Tag) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("ModifiedOn", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
