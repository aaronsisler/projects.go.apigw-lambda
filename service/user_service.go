package service

import (
	"context"
	"errors"

	"github.com/aaronsisler/projects.go.apigw-lambda/dao"
	model "github.com/aaronsisler/projects.go.apigw-lambda/models"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (s *UserService) GetUserById(ctx context.Context, userId string) (*model.User, error) {
	if userId == "" {
		return nil, errors.New("userId cannot be empty")
	}

	return s.userDao.GetUserById(ctx, userId)
}
