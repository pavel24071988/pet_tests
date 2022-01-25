package tests_test

import (
	"github.com/golang/mock/gomock"
	"pet_tests/pkg/service"
	"pet_tests/tests/mocks"
	"testing"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser := &service.User{IUser: mockUser}

	mockUser.EXPECT().AddUser(1, "user created").Return(nil).Times(1)
	testUser.Use()
}
