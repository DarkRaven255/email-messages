package repository

import (
	"email-messages/domain"
	"email-messages/domain/domainmodel"

	"github.com/gocql/gocql"
)

type repository struct {
	db *gocql.Session
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

func NewEntryRepository(dbConn *gocql.Session) domain.MessagesRepository {

	return &repository{
		db: dbConn,
	}
}
