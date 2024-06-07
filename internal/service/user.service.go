package service

import "github.com/duonglc27092000/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo:  repo.NewUserRepo(),
	}
}

// user service - us
func (us *UserService) GetInfoUserService() string {
	return us.userRepo.GetInfoUserRepo()
}
