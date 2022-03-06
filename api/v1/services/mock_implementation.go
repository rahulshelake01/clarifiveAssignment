package services

import (
	"clarifiveassignment/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type mockFamilyRepository struct {
	m mock.Mock
}

func (m *mockFamilyRepository) Add(request *models.AddMemberRequest) error {

	fmt.Printf("Mock Add() repository method is called with request : %+v\n", request)
	mockArgs := m.m.Called(request)
	return mockArgs.Error(0)

}

func (m *mockFamilyRepository) Get(request *models.GetMemberRequest) (*[]models.Member, error) {
	fmt.Printf("Mock Get() repository method is called with email : %+v\n", request)
	mockArgs := m.m.Called(request)
	return mockArgs.Get(0).(*[]models.Member), mockArgs.Error(1)
}
