package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BlogArticleMgr struct {
	*_BaseMgr
}

// BlogArticleMgr open func
func BlogArticleMgr(db *gorm.DB) *_BlogArticleMgr {
	if db == nil {
		panic(fmt.Errorf("BlogArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BlogArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Model(BlogArticle{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BlogArticleMgr) GetTableName() string {
	return "blog_blog_article"
}

// Get 获取
func (obj *_BlogArticleMgr) Get() (result BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BlogArticleMgr) Gets() (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BlogArticleMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTagID tag_id获取 标签ID
func (obj *_BlogArticleMgr) WithTagID(tagID uint) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithTitle title获取 文章标题
func (obj *_BlogArticleMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDesc desc获取 简述
func (obj *_BlogArticleMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithContent content获取
func (obj *_BlogArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreatedOn created_on获取
func (obj *_BlogArticleMgr) WithCreatedOn(createdOn int) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithCreatedBy created_by获取 创建人
func (obj *_BlogArticleMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithModifiedOn modified_on获取 修改时间
func (obj *_BlogArticleMgr) WithModifiedOn(modifiedOn uint) Option {
	return optionFunc(func(o *options) { o.query["modified_on"] = modifiedOn })
}

// WithModifiedBy modified_by获取 修改人
func (obj *_BlogArticleMgr) WithModifiedBy(modifiedBy string) Option {
	return optionFunc(func(o *options) { o.query["modified_by"] = modifiedBy })
}

// WithDeletedOn deleted_on获取
func (obj *_BlogArticleMgr) WithDeletedOn(deletedOn uint) Option {
	return optionFunc(func(o *options) { o.query["deleted_on"] = deletedOn })
}

// WithState state获取 状态 0为禁用1为启用
func (obj *_BlogArticleMgr) WithState(state uint8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// GetByOption 功能选项模式获取
func (obj *_BlogArticleMgr) GetByOption(opts ...Option) (result BlogArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BlogArticleMgr) GetByOptions(opts ...Option) (results []*BlogArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_BlogArticleMgr) GetFromID(id uint) (result BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BlogArticleMgr) GetBatchFromID(ids []uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTagID 通过tag_id获取内容 标签ID
func (obj *_BlogArticleMgr) GetFromTagID(tagID uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`tag_id` = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量查找 标签ID
func (obj *_BlogArticleMgr) GetBatchFromTagID(tagIDs []uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 文章标题
func (obj *_BlogArticleMgr) GetFromTitle(title string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 文章标题
func (obj *_BlogArticleMgr) GetBatchFromTitle(titles []string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 简述
func (obj *_BlogArticleMgr) GetFromDesc(desc string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 简述
func (obj *_BlogArticleMgr) GetBatchFromDesc(descs []string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_BlogArticleMgr) GetFromContent(content string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找
func (obj *_BlogArticleMgr) GetBatchFromContent(contents []string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_BlogArticleMgr) GetFromCreatedOn(createdOn int) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_on` = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_BlogArticleMgr) GetBatchFromCreatedOn(createdOns []int) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_on` IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建人
func (obj *_BlogArticleMgr) GetFromCreatedBy(createdBy string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建人
func (obj *_BlogArticleMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromModifiedOn 通过modified_on获取内容 修改时间
func (obj *_BlogArticleMgr) GetFromModifiedOn(modifiedOn uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_on` = ?", modifiedOn).Find(&results).Error

	return
}

// GetBatchFromModifiedOn 批量查找 修改时间
func (obj *_BlogArticleMgr) GetBatchFromModifiedOn(modifiedOns []uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_on` IN (?)", modifiedOns).Find(&results).Error

	return
}

// GetFromModifiedBy 通过modified_by获取内容 修改人
func (obj *_BlogArticleMgr) GetFromModifiedBy(modifiedBy string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_by` = ?", modifiedBy).Find(&results).Error

	return
}

// GetBatchFromModifiedBy 批量查找 修改人
func (obj *_BlogArticleMgr) GetBatchFromModifiedBy(modifiedBys []string) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_by` IN (?)", modifiedBys).Find(&results).Error

	return
}

// GetFromDeletedOn 通过deleted_on获取内容
func (obj *_BlogArticleMgr) GetFromDeletedOn(deletedOn uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`deleted_on` = ?", deletedOn).Find(&results).Error

	return
}

// GetBatchFromDeletedOn 批量查找
func (obj *_BlogArticleMgr) GetBatchFromDeletedOn(deletedOns []uint) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`deleted_on` IN (?)", deletedOns).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态 0为禁用1为启用
func (obj *_BlogArticleMgr) GetFromState(state uint8) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态 0为禁用1为启用
func (obj *_BlogArticleMgr) GetBatchFromState(states []uint8) (results []*BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`state` IN (?)", states).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BlogArticleMgr) FetchByPrimaryKey(id uint) (result BlogArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
