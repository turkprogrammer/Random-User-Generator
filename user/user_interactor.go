package user

type UserInteractor struct {
	userRepository UserRepository
}

func (i *UserInteractor) GetRandomUser() (*User, error) {
	return i.userRepository.GetRandomUser()
}

func NewRandomUserInteractor(repo UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: repo,
	}
}
