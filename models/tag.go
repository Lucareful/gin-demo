package models

import (
	"github.com/luenci/errors"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
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

// TagList Tag 列表
type TagList struct {
	TotalCount int64
	Item       []*Tag
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

// ListTag 获取所有 tag 分页返回
func (tl *TagList) ListTag(r *paginate.Query) error {
	fieldList, err := getFieldList(GetDB(), new(Tag))
	if err != nil {
		return err
	}
	err = GetDB().Scopes(
		paginate.Paginate(r, nil, fieldList),
	).Find(&tl.Item).Limit(-1).Offset(-1).Count(&tl.TotalCount).Error
	return errors.WithStack(err)
}

func NewTag() *Tag {
	return &Tag{}
}

func NewListTag() *TagList {
	return &TagList{}
}

// getFieldList 获取指定数据表的所有字段
func getFieldList(db *gorm.DB, bean interface{}) ([]string, error) {
	var fieldList []string
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(bean)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for i := range stmt.Schema.Fields {
		fieldName := stmt.Schema.Fields[i].DBName
		if fieldName != "" {
			fieldList = append(fieldList, fieldName)
		}
	}
	return fieldList, nil
}
