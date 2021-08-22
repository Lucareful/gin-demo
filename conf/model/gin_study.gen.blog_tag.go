package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BlogTagMgr struct {
	*_BaseMgr
}

// BlogTagMgr open func
func BlogTagMgr(db *gorm.DB) *_BlogTagMgr {
	if db == nil {
		panic(fmt.Errorf("BlogTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BlogTagMgr{_BaseMgr: &_BaseMgr{DB: db.Model(BlogTag{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BlogTagMgr) GetTableName() string {
	return "blog_blog_tag"
}

// Get 获取
func (obj *_BlogTagMgr) Get() (result BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BlogTagMgr) Gets() (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BlogTagMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 标签名称
func (obj *_BlogTagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreatedOn created_on获取 创建时间
func (obj *_BlogTagMgr) WithCreatedOn(createdOn uint) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithCreatedBy created_by获取 创建人
func (obj *_BlogTagMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithModifiedOn modified_on获取 修改时间
func (obj *_BlogTagMgr) WithModifiedOn(modifiedOn uint) Option {
	return optionFunc(func(o *options) { o.query["modified_on"] = modifiedOn })
}

// WithModifiedBy modified_by获取 修改人
func (obj *_BlogTagMgr) WithModifiedBy(modifiedBy string) Option {
	return optionFunc(func(o *options) { o.query["modified_by"] = modifiedBy })
}

// WithDeletedOn deleted_on获取
func (obj *_BlogTagMgr) WithDeletedOn(deletedOn uint) Option {
	return optionFunc(func(o *options) { o.query["deleted_on"] = deletedOn })
}

// WithState state获取 状态 0为禁用、1为启用
func (obj *_BlogTagMgr) WithState(state uint8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// GetByOption 功能选项模式获取
func (obj *_BlogTagMgr) GetByOption(opts ...Option) (result BlogTag, err error) {
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
func (obj *_BlogTagMgr) GetByOptions(opts ...Option) (results []*BlogTag, err error) {
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
func (obj *_BlogTagMgr) GetFromID(id uint) (result BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BlogTagMgr) GetBatchFromID(ids []uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 标签名称
func (obj *_BlogTagMgr) GetFromName(name string) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 标签名称
func (obj *_BlogTagMgr) GetBatchFromName(names []string) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容 创建时间
func (obj *_BlogTagMgr) GetFromCreatedOn(createdOn uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_on` = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量查找 创建时间
func (obj *_BlogTagMgr) GetBatchFromCreatedOn(createdOns []uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_on` IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建人
func (obj *_BlogTagMgr) GetFromCreatedBy(createdBy string) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建人
func (obj *_BlogTagMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromModifiedOn 通过modified_on获取内容 修改时间
func (obj *_BlogTagMgr) GetFromModifiedOn(modifiedOn uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_on` = ?", modifiedOn).Find(&results).Error

	return
}

// GetBatchFromModifiedOn 批量查找 修改时间
func (obj *_BlogTagMgr) GetBatchFromModifiedOn(modifiedOns []uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_on` IN (?)", modifiedOns).Find(&results).Error

	return
}

// GetFromModifiedBy 通过modified_by获取内容 修改人
func (obj *_BlogTagMgr) GetFromModifiedBy(modifiedBy string) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_by` = ?", modifiedBy).Find(&results).Error

	return
}

// GetBatchFromModifiedBy 批量查找 修改人
func (obj *_BlogTagMgr) GetBatchFromModifiedBy(modifiedBys []string) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`modified_by` IN (?)", modifiedBys).Find(&results).Error

	return
}

// GetFromDeletedOn 通过deleted_on获取内容
func (obj *_BlogTagMgr) GetFromDeletedOn(deletedOn uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`deleted_on` = ?", deletedOn).Find(&results).Error

	return
}

// GetBatchFromDeletedOn 批量查找
func (obj *_BlogTagMgr) GetBatchFromDeletedOn(deletedOns []uint) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`deleted_on` IN (?)", deletedOns).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态 0为禁用、1为启用
func (obj *_BlogTagMgr) GetFromState(state uint8) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态 0为禁用、1为启用
func (obj *_BlogTagMgr) GetBatchFromState(states []uint8) (results []*BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`state` IN (?)", states).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BlogTagMgr) FetchByPrimaryKey(id uint) (result BlogTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
