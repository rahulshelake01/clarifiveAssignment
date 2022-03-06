package services

import (
	"clarifiveassignment/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type familySuit struct {
	suite.Suite
	mockFamilyRepository *mockFamilyRepository
	familyService        FamilyServiceInterface
}

func TestMemberTestSuite(t *testing.T) {
	suite.Run(t, new(familySuit))
}

func (u *familySuit) SetupTest() {

	u.mockFamilyRepository = new(mockFamilyRepository)
	u.familyService = FamilyService(u.mockFamilyRepository)

}

func (u *familySuit) TestAddMemberWithFail() {

	var (
		request          = getAddMemberRequest()
		expectedResponse = getAddMemberFailResponse()
	)
	u.mockFamilyRepository.m.On("Add", mock.Anything).Return(errors.New("Failed to login"))
	response := u.familyService.Add(request)
	u.Equal(expectedResponse, response)

	u.mockFamilyRepository.m.AssertExpectations(u.Suite.T())

}

func (u *familySuit) TestAddMemberWithSuccess() {

	var (
		request          = getAddMemberRequest()
		expectedResponse = getAddMemberResponse()
	)

	u.mockFamilyRepository.m.On("Add", mock.Anything).Return(nil)

	response := u.familyService.Add(request)
	u.Equal(expectedResponse, response)

	u.mockFamilyRepository.m.AssertExpectations(u.Suite.T())

}

func (u *familySuit) TestGetMemberWithFail() {

	var (
		request          = getMemberRequest()
		expectedResponse = getMemberFailResponse()
	)

	u.mockFamilyRepository.m.On("Get", mock.Anything).Return(&[]models.Member{}, errors.New("Error"))

	response := u.familyService.Get(request)
	u.Equal(expectedResponse, response)

	u.mockFamilyRepository.m.AssertExpectations(u.Suite.T())

}

func (u *familySuit) TestGetMemberWithSuccess() {

	var (
		request          = getMemberRequest()
		expectedResponse = getMemberSuccessResponse()
	)

	u.mockFamilyRepository.m.On("Get", mock.Anything).Return(&[]models.Member{}, nil)

	response := u.familyService.Get(request)
	u.Equal(expectedResponse, response)

	u.mockFamilyRepository.m.AssertExpectations(u.Suite.T())

}

func getAddMemberRequest() *models.AddMemberRequest {

	return &models.AddMemberRequest{
		FirstName: "RahulTest",
		LastName:  "ShelakeTest",
	}
}

func getAddMemberFailResponse() *models.AddMemberResponse {
	return &models.AddMemberResponse{
		Success:    false,
		StatusCode: 300,
		Message:    "Something went wrong.",
	}
}

func getAddMemberResponse() *models.AddMemberResponse {
	return &models.AddMemberResponse{
		Success:    true,
		StatusCode: 200,
		Message:    "Member added successfully.",
	}
}

func getMemberRequest() *models.GetMemberRequest {
	return &models.GetMemberRequest{
		Name: "RahulTest",
	}
}

func getMemberFailResponse() *models.GetMemberResponse {
	return &models.GetMemberResponse{
		Success:    false,
		StatusCode: 300,
		Message:    "Something went wrong.",
	}
}

func getMemberSuccessResponse() *models.GetMemberResponse {
	return &models.GetMemberResponse{
		Success:    true,
		StatusCode: 200,
		Message:    "Member list",
		Data:       []models.Member{},
	}
}
