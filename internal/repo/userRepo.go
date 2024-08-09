package repo

// type UserRepo struct {
// }

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "masterbi"
// }

// Interface Version
type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct {
}

// GetUserByEmail implements IUserRepo.
func (u *userRepo) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
