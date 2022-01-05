package service

import (
	"github.com/luenci/go-gin-example/models"
	"github.com/luenci/go-gin-example/pkg/http/paginate"
	"github.com/luenci/go-gin-example/types/request"
)

type userService struct{}

func (u *userService) ListUserService(query *paginate.Query) (*models.UserList, error) {
	userTag := models.NewListUser()

	err := userTag.ListUser(query)
	if err != nil {
		return nil, err
	}

	return userTag, nil
}

func (u *userService) GetUserService(id uint) (*models.User, error) {
	user := models.NewUser()

	if err := user.GetUserById(id); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateUserService(r request.CreateUserRequest) (*models.User, error) {

	user := &models.User{
		UserName: r.UserName,
		Password: r.Password,
	}

	if err := user.CreateUser(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateUserService(r request.UpdateUserRequest) (*models.User, error) {
	user := &models.User{
		Id:       r.Id,
		UserName: r.UserName,
		Password: r.Password,
	}

	if err := user.UpdateUserByID(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteUserService(id uint) (*models.User, error) {
	user := models.NewUser()

	if err := user.DeleteUserById(id); err != nil {
		return nil, err
	}

	return user, nil

}
