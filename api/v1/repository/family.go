package repository

import (
	"clarifiveassignment/models"
	"time"

	"github.com/jinzhu/gorm"
)

type FamilyRepositoryInterface interface {
	Add(request *models.AddMemberRequest) error
	Get(request *models.GetMemberRequest) (*[]models.Member, error)
}

type FamilyRepositoryStruct struct {
	DB *gorm.DB
}

func FamilyRepository(DB *gorm.DB) FamilyRepositoryInterface {
	return FamilyRepositoryStruct{DB}
}

func (repo FamilyRepositoryStruct) Add(request *models.AddMemberRequest) error {

	member := models.Member{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Father:    request.Father,
		Mother:    request.Mother,
		Brother:   request.Brother,
		Sister:    request.Sister,
		Uncle:     request.Uncle,
		Aunt:      request.Aunt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	op := repo.DB.Create(member)
	return op.Error
}

func (repo FamilyRepositoryStruct) Get(request *models.GetMemberRequest) (*[]models.Member, error) {

	var members = new([]models.Member)
	op := repo.DB.Find(&members, "first_name = ?", request.Name)
	return members, op.Error
}
