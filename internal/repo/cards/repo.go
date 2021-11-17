package repo_cards

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/pay-card-api/internal/model"
)

// Repo is DAO for Template
type Repo interface {
	Add(card *model.Card) (uint64, error)
	Get(cardID uint64) (*model.Card, error)
	List(limit uint64, cursor uint64) ([]model.Card, error)
	Remove(cardID uint64) (bool, error)
}

type repo struct {
	db *sqlx.DB
}

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewCardRepo(db *sqlx.DB) Repo {
	return &repo{db: db}
}

func (c repo) Add(card *model.Card) (uint64, error) {
	query := psql.Insert("cards").
		Columns("owner_id, payment_system, number, holder_name, expiration_date, cvccvv").
		Values(card.OwnerId, card.PaymentSystem, card.Number, card.HolderName, card.ExpirationDate, card.CvcCvv).
		Suffix("RETURNING id").
		RunWith(c.db)

	rows, err := query.Query()

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	} else {
		return 0, sql.ErrNoRows
	}
}

func (c repo) Get(cardID uint64) (*model.Card, error) {
	query := psql.Select("*").
		From("cards").
		Where(sq.Eq{"id": cardID}).
		RunWith(c.db)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	res := &model.Card{}
	err = c.db.Select(&res, sql, args...)
	return res, err
}

func (c repo) List(limit uint64, cursor uint64) ([]model.Card, error) {
	query := psql.Select("*").
		From("cards").
		Offset(cursor).
		Limit(limit).
		OrderBy("id ASC").
		RunWith(c.db)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []model.Card
	err = c.db.Select(&res, sql, args...)
	return res, err
}

func (c repo) Remove(cardID uint64) (bool, error) {
	query := psql.Delete("card").
		Where(sq.Eq{"id": cardID}).
		RunWith(c.db)

	_, err := query.Exec()
	if err != nil {
		return false, err
	}

	return true, nil
}
