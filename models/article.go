package models

import (
	"github.com/luenci/errors"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ArticleList Article 列表
type ArticleList struct {
	TotalCount int64
	Item       []*Article
}

// GetArticle 获取单个 Article 记录.
func (t *Article) GetArticle(id uint) error {
	err := GetDB().Where("id = ?", id).First(&t).Error
	return errors.WithStack(err)
}

// UpdateArticle 更新 Article 记录.
func (t *Article) UpdateArticle(r request.UpdateArticleRequest) error {
	err := GetDB().Model(&Article{}).Where("id = ?", r.ID).Updates(&Article{
		Title: r.Title,
		State: r.State,
	}).First(&t).Error
	return errors.WithStack(err)
}

// DeleteArticle 删除 Article 记录.
func (t *Article) DeleteArticle(id uint) error {
	err := GetDB().Delete(&Article{}, id).First(&t).Error
	return errors.WithStack(err)
}

// ExistArticleByName 判断 Article 记录是否存在.
func (t *Article) ExistArticleByName(name string) error {
	err := GetDB().Select("id").Where("name = ?", name).First(&t).Error
	return errors.WithStack(err)
}

// CreateArticle 为 Article 表增加一条记录.
func (t *Article) CreateArticle(r request.CreateArticleRequest) error {
	err := GetDB().Create(&Article{
		Title:     r.Title,
		State:     r.State,
		CreatedBy: r.CreatedBy}).First(&t).Error
	return errors.WithStack(err)
}

// ListArticle 获取所有 Article 分页返回
func (tl *ArticleList) ListArticle(r *paginate.Query) error {
	fieldList, err := getFieldList(GetDB(), new(Article))
	if err != nil {
		return err
	}
	err = GetDB().Scopes(
		paginate.Paginate(r, nil, fieldList),
	).Find(&tl.Item).Limit(-1).Offset(-1).Count(&tl.TotalCount).Error
	return errors.WithStack(err)
}

func NewArticle() *Article {
	return &Article{}
}

func NewListArticle() *ArticleList {
	return &ArticleList{}
}
