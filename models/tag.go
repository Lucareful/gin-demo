package models

import (
	"github.com/luenci/go-gin-example/types/request"
	"github.com/marmotedu/errors"
	"gorm.io/gorm"
)

// Tag tag 表的结构体
type Tag struct {
	gorm.Model

	Name       string `json:"name" gorm:"unique_index:name"`
	CreatedBy  string `json:"created_by" gorm:"unique_index:name"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state" gorm:"unique_index:name"`
}

// GetTag 获取单个 tag 记录.
func (t *Tag) GetTag(id uint) error {
	err := GetDB().Where("id = ?", id).First(&t).Error
	return errors.WithStack(err)
}

// UpdateTag 更新 tag 记录.
func (t *Tag) UpdateTag(r request.UpdateTagRequest) error {
	err := GetDB().Model(&Tag{}).Where("id = ?", r.ID).Updates(&Tag{
		Name:  r.Name,
		State: r.State,
	}).First(&t).Error
	return errors.WithStack(err)
}

// DeleteTag 删除 tag 记录.
func (t *Tag) DeleteTag(id uint) error {
	err := GetDB().Delete(&Tag{}, id).First(&t).Error
	return errors.WithStack(err)
}

// ExistTagByName 判断 tag 记录是否存在.
func (t *Tag) ExistTagByName(name string) error {
	err := GetDB().Select("id").Where("name = ?", name).First(&t).Error
	return errors.WithStack(err)
}

// CreateTag 为 tag 表增加一条记录.
func (t *Tag) CreateTag(r request.CreateTagRequest) error {
	err := GetDB().Create(&Tag{
		Name:      r.Name,
		State:     r.State,
		CreatedBy: r.CreatedBy}).First(&t).Error
	return errors.WithStack(err)
}

func NewTag() *Tag {
	return &Tag{}
}
