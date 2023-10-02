package usecase

import (
	"beautify/pkg/domain"
	mock "beautify/pkg/mock/userRepoMock"
	"beautify/pkg/utils/response"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctrl)

	userUsecase := NewUserUseCase(mockUserRepository)

	ctx := context.Background()

	// Test data
	user := domain.Users{
		UserName:  "JerinMathai",
		FirstName: "Jerin",
		LastName:  "Mathai",
		Age:       25,
		Email:     "jerin@gmail.com",
		Phone:     "9446371883",
		Password:  "password",
	}

	// Test case : 1 "Success"
	t.Run("Success signup", func(t *testing.T) {
		mockUserRepository.EXPECT().FindUser(ctx, user).Return(domain.Users{}, nil)
		mockUserRepository.EXPECT().SaveUser(ctx, gomock.Any()).Return(response.UserSignUp{}, nil)

		_, err := userUsecase.SignUp(ctx, user)

		assert.NoError(t, err)
	})

	// Test case : 2 "User already exists"
	existingUser := user
	existingUser.ID = 1
	t.Run("User already exists, should return an error", func(t *testing.T) {
		mockUserRepository.EXPECT().FindUser(ctx, user).Return(existingUser, nil)

		_, err := userUsecase.SignUp(ctx, user)

		assert.EqualError(t, err, "JerinMathai user already exists")
	})

	// Test case : 1 "Failed to save"
	// Passing null value to produce database error
	// user.Email = ""
	// t.Run("Error saving user, should return the error", func(t *testing.T) {
	// 	expectedErr := errors.New("failed to save user")
	// 	mockUserRepository.EXPECT().FindUser(ctx, user).Return(domain.Users{}, nil)
	// 	fmt.Println("------------------------", user)
	// 	mockAuthRepository.EXPECT().SaveUser(ctx, user).Return(expectedErr)

	// 	err := authUsecase.SignUp(ctx, user)
	// 	fmt.Println("----------------", err)
	// 	assert.EqualError(t, err, expectedErr.Error())
	// })

}
