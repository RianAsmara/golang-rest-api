package services

import (
	"kulina-interview-test/domain/dto"
	"kulina-interview-test/domain/model"
	"kulina-interview-test/domain/repository"

	"gorm.io/gorm"
)

type supplierService struct {
	supplierRepo repository.SupplierRepository
}

func (service *supplierService) CreateStore(request dto.CreateStoreRequest) dto.CreateStoreResponse {
	store := model.Supplier{
		Store: model.Store{
			FullAddress: request.FullAddress,
			Latitude:    request.Latitude,
			Longitude:   request.Longitude,
		},
	}
	result, err := service.supplierRepo.CreateStore(store)

	if err != nil {
		return dto.CreateStoreResponse{}
	}

	response := dto.CreateStoreResponse{
		ID:          result,
		FullAddress: request.FullAddress,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
	}
	return response
}

type SupplierService interface {
	CreateStore(request dto.CreateStoreRequest) dto.CreateStoreResponse
}

func NewSupplierService(db *gorm.DB) SupplierService {
	return &supplierService{supplierRepo: repository.NewSupplierRepository(db)}
}
