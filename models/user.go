package models

import (
	"github.com/luenci/errors"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
)

type User struct {
	Id       uint   `json:"id" gorm:"primarykey"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// UserList user 列表.
type UserList struct {
	TotalCount int64
	Item       []*Tag
}

// TableName 返回对应的数据库表名.
func (u *User) TableName() string {
	return "users"
}

// CreateUser 创建用户.
func (u *User) CreateUser() error {
	err := GetDB().Create(&u).Error
	return errors.WithStack(err)
}

// GetUserById 通过 ID 获取用户信息.
func (u *User) GetUserById(id uint) error {
	err := GetDB().First(&u, id).Error
	return errors.WithStack(err)
}

// UpdateUserByID 更新用户信息.
func (u *User) UpdateUserByID() error {
	err := GetDB().Select("id = ?", u.Id).Updates(&u).First(&u).Error
	return errors.WithStack(err)
}

// DeleteUserById 通过 ID 删除用户.
func (u *User) DeleteUserById(id uint) error {
	err := GetDB().Delete(&User{}, id).First(&u).Error
	return errors.WithStack(err)
}

// ListUser 获取用户列表.
func (ul *UserList) ListUser(r *paginate.Query) error {
	fieldList, err := getFieldList(GetDB(), new(Tag))
	if err != nil {
		return err
	}
	err = GetDB().Scopes(
		paginate.Paginate(r, nil, fieldList),
	).Find(&ul.Item).Limit(-1).Offset(-1).Count(&ul.TotalCount).Error
	return errors.WithStack(err)
}

func NewListUser() *UserList {
	return &UserList{}
}

func NewUser() *User {
	return &User{}
}
