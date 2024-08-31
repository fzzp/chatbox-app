package dao

import "chatbox-app/models"

type UserDao struct{}

func (*UserDao) CreateUser(user *models.User) {
	Create(user)
}

func (*UserDao) GetById(id uint) *models.User {
	return QueryRow(&models.User{}, "id = ?", id)
}

func (*UserDao) GetByEmail(email string) *models.User {
	return QueryRow(&models.User{}, "email = ?", email)
}
