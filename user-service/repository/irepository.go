package repository

import "user-service/domain"

type IUserServiceRepository interface {
	GetAllUsers(pageNum, pageSize int) ([]domain.User, error)
	GetUserByID(id int64) (domain.User, error)
	CreateUser(name string) (domain.User, error)
}
