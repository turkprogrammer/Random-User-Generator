// user/user_interactor.go
package user

type UserInteractor struct {
	userRepository UserRepository
}

func NewUserInteractor(userRepository UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

func (i *UserInteractor) GetRandomUser() (*User, error) {
	return i.userRepository.GetRandomUser()
}

// Используйте UserRepository вместо RandomUserRepository
func NewRandomUserInteractor(repo UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: repo,
	}
}
