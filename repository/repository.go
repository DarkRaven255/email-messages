package repository

import (
	"email-messages/domain"
	"email-messages/domain/domainmodel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(entry *domainmodel.Model) error {
	return nil
}

func (r *repository) Delete() error {

	return nil
}

func (r *repository) GetByMagicNumber(magicNumber *string) ([]*domainmodel.Model, error) {
	return nil, nil
}

func NewEntryRepository(dbConn *gorm.DB) domain.MessagesRepository {

	return &repository{
		db: dbConn,
	}
}
