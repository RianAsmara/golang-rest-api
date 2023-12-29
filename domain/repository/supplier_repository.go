package repository

import (
	"kulina-interview-test/domain/model"

	"gorm.io/gorm"
)

type dbSupplier struct {
	Conn *gorm.DB
}

func (db *dbSupplier) CreateStore(store model.Store) (model.Store, error) {
	db.Conn.Create(&store)

	return store, nil
}

type SupplierRepository interface {
	CreateStore(store model.Store) (model.Store, error)
}

func NewSupplierRepository(Conn *gorm.DB) SupplierRepository {
	return &dbSupplier{Conn: Conn}
}
