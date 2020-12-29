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

func (r *repository) GetByMagicNumber(magicNumber *int) (*[]domainmodel.Model, error) {
	var (
		model  domainmodel.Model
		models []domainmodel.Model
		q      = `SELECT id, email, title, content from em.messages WHERE magic_number = ? ALLOW FILTERING`
	)

	iter := r.session.Query(q, magicNumber).Iter()
	for iter.Scan(&model.Id, &model.Email, &model.Title, &model.Content) {
		models = append(models, model)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &models, nil
}

func (r *repository) GetByEmail(email *string) (*[]domainmodel.Model, error) {
	var (
		model  domainmodel.Model
		models []domainmodel.Model
		q      = `SELECT id, timestamp, email, title, content, magic_number from em.messages WHERE email = ? ALLOW FILTERING`
	)

	iter := r.session.Query(q, email).Iter()
	for iter.Scan(&model.Id, &model.Timestamp, &model.Email, &model.Title, &model.Content, &model.MagicNumber) {
		models = append(models, model)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &models, nil
}

func (r *repository) Delete(id *gocql.UUID) error {
	var (
		q = `DELETE FROM em.messages WHERE id = ?`
	)

	err := r.session.Query(q, id).Exec()
	if err != nil {
		return err
	}

	return nil
}

func NewEntryRepository(dbConn *gocql.Session) domain.MessagesRepository {

	return &repository{
		session: dbConn,
	}
}
