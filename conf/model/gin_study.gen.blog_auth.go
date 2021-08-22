package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BlogAuthMgr struct {
	*_BaseMgr
}

// BlogAuthMgr open func
func BlogAuthMgr(db *gorm.DB) *_BlogAuthMgr {
	if db == nil {
		panic(fmt.Errorf("BlogAuthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BlogAuthMgr{_BaseMgr: &_BaseMgr{DB: db.Model(BlogAuth{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BlogAuthMgr) GetTableName() string {
	return "blog_blog_auth"
}

// Get 获取
func (obj *_BlogAuthMgr) Get() (result BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BlogAuthMgr) Gets() (results []*BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BlogAuthMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 账号
func (obj *_BlogAuthMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码
func (obj *_BlogAuthMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_BlogAuthMgr) GetByOption(opts ...Option) (result BlogAuth, err error) {
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
func (obj *_BlogAuthMgr) GetByOptions(opts ...Option) (results []*BlogAuth, err error) {
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
func (obj *_BlogAuthMgr) GetFromID(id uint) (result BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BlogAuthMgr) GetBatchFromID(ids []uint) (results []*BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 账号
func (obj *_BlogAuthMgr) GetFromUsername(username string) (results []*BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 账号
func (obj *_BlogAuthMgr) GetBatchFromUsername(usernames []string) (results []*BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_BlogAuthMgr) GetFromPassword(password string) (results []*BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_BlogAuthMgr) GetBatchFromPassword(passwords []string) (results []*BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BlogAuthMgr) FetchByPrimaryKey(id uint) (result BlogAuth, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
