package user

import "context"

// verity that the repository implements the interface
type MockUserRepository struct{}

func CreateMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) CreateUser(
	ctx context.Context,
) (*UserEntity, error) {
	return &UserEntity{
		Id:       1,
		Email:    "test@email.com",
		Password: "password",
		Name:     "name",
		Gender:   "male",
		Age:      22,
	}, nil
}

func (m *MockUserRepository) GetUserByEmailAndPassword(
	ctx context.Context,
	criteria *EmailAndPasswordCriteria,
) (*UserEntity, error) {
	return &UserEntity{
		Id:       1,
		Email:    criteria.Email,
		Password: criteria.Password,
		Name:     "name",
		Gender:   "male",
		Age:      22,
	}, nil
}
