package tests_test

import (
	"github.com/golang/mock/gomock"
	"pet_tests/internal/service"
	"pet_tests/pkg/tests/mocks"
	"testing"
)

func TestAdd(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockProduct := mocks.NewMockIProductInterface(mockCtrl)
	testProduct := &service.Product{IProduct: mockProduct}

	cost := float32(1500.5)

	mockProduct.EXPECT().AddProduct("bolgarka", "stihl", cost).Return(1)
	testProduct.Add()
}

func TestUpdate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProduct := mocks.NewMockIProductInterface(mockController)
	testProduct := service.Product{IProduct: mockProduct}

	cost := float32(3000.8)

	mockProduct.EXPECT().UpdateProduct(1, "usm", "stihl", cost)
	testProduct.Update()
}

func TestDelete(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProduct := mocks.NewMockIProductInterface(mockController)
	testProduct := service.Product{IProduct: mockProduct}

	mockProduct.EXPECT().DeleteProduct(1)
	testProduct.Delete()
}
