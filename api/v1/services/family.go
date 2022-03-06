package services

import (
	"clarifiveassignment/api/v1/repository"
	"clarifiveassignment/models"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("MY_JWT_SECRET")

type Claims struct {
	Uid   int64  `json:"uid"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type FamilyServiceInterface interface {
	Add(request *models.AddMemberRequest) *models.AddMemberResponse
	Get(request *models.GetMemberRequest) *models.GetMemberResponse
}

type FamilyServiceStruct struct {
	FamilyRepository repository.FamilyRepositoryInterface
}

func FamilyService(familyRepo repository.FamilyRepositoryInterface) FamilyServiceInterface {
	return FamilyServiceStruct{FamilyRepository: familyRepo}
}

func (familyService FamilyServiceStruct) Add(request *models.AddMemberRequest) *models.AddMemberResponse {

	var (
		response = new(models.AddMemberResponse)
		err      error
	)

	if err != nil {
		fmt.Println("Failed to hash password.")
		response.Message = "Something went wrong."
		return response
	}

	err = familyService.FamilyRepository.Add(request)

	if err != nil {
		fmt.Println("Failed to add member.")
		response.Message = "Something went wrong."
		response.Success = false
		response.StatusCode = 300
		return response
	}

	response.Message = "Member added successfully."
	response.Success = true
	response.StatusCode = 200
	return response
}

func (familyService FamilyServiceStruct) Get(request *models.GetMemberRequest) *models.GetMemberResponse {

	var response = new(models.GetMemberResponse)

	members, err := familyService.FamilyRepository.Get(request)

	if err != nil {
		fmt.Println("Failed to fetch members.")
		response.Message = "Something went wrong."
		response.Success = false
		response.StatusCode = 300
		return response
	}

	response.Success = true
	response.StatusCode = 200
	response.Message = "Member list"
	response.Data = *members
	return response
}
