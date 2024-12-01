package services

import (
    "github.com/gglzc/fishMachine/models/user"
    "github.com/gglzc/fishMachine/repository/user"
)

type UserService struct {
    UserRepo *repositories.UserRepository
}

func (us *UserService) GetUserByID(id int64) (*models.User, error) {
    return us.UserRepo.GetUserByID(id)
}
func (us *UserService) CreateUser(user models.User) (error) {
    err:=us.UserRepo.CreateUser(user)
	if(err!=nil){
		return err
	}
	return nil
}

func (us *UserService) UpdateBalance(userID int64, amount int64) error {
    return us.UserRepo.UpdateBalance(userID, amount)
}

