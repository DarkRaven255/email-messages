package repository

import (
	"email-messages/domain"
	"email-messages/domain/domainmodel"

	"github.com/gocql/gocql"
)

type repository struct {
	session *gocql.Session
}

func (r *repository) Create(entry *domainmodel.Model) error {
	q := `
	INSERT INTO em.messages (
		id,
		timestamp,
		email,
		title,
		content,
		magic_number
	)
	VALUES (uuid(), toTimestamp(now()), ?, ?, ?, ?) USING TTL 300
	`
	err := r.session.Query(q,
		entry.Email,
		entry.Title,
		entry.Content,
		entry.MagicNumber).Exec()
	if err != nil {
		return err
	}

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
		session: dbConn,
	}
}
