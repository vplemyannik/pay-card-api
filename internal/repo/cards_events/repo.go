package repo_cards_events

import (
	"context"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/pay-card-api/internal/model"
	"github.com/ozonmp/pay-card-api/internal/pkg/db"
	"github.com/pkg/errors"
	"time"
)

import (
	. "github.com/ozonmp/pay-card-api/internal/model"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Repo interface {
	Lock(n uint64) ([]model.CardEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.CardEvent) error
	Remove(eventIDs []uint64) error
}

type repo struct {
	db *sqlx.DB
}

func NewCardEventsRepo(db *sqlx.DB) Repo {
	return &repo{db: db}
}

const (
	lockID   int = 2 ^ 32 - 1
	entityID int = 2 ^ 32 - 2
)

func (r repo) Lock(n uint64) ([]model.CardEvent, error) {

	var events []model.CardEvent

	ctx := context.Background()
	err := db.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) error {
		isAcquired, err := db.AcquireTryLock(ctx, tx, int32(lockID), int32(entityID))
		if !isAcquired || err != nil {
			return errors.Wrap(err, "database.WithTx")
		}

		query := psql.Select("*").
			From("cards_events").
			Where("status = 'New'").
			Limit(n).
			OrderBy("updated_at ASC").
			RunWith(r.db)

		sql, args, err := query.ToSql()
		if err != nil {
			return err
		}

		var res []model.CardEvent
		err = r.db.Select(&res, sql, args...)
		if err != nil {
			return err
		}

		ids := make([]uint64, 0, len(res))
		for _, ev := range res {
			ids = append(ids, ev.ID)
		}

		updateQuery := psql.Update("cards_events").
			Set("status", Locked).
			Where(sq.Eq{"id": ids}).
			RunWith(r.db)

		_, err = updateQuery.Exec()

		events = res

		return nil
	})

	return events, err
}

func (r repo) Unlock(eventIDs []uint64) error {
	updateQuery := psql.Update("cards_events").
		Set("status", Processed).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": eventIDs}).
		RunWith(r.db)

	_, err := updateQuery.Exec()

	return err
}

func (r repo) Add(events []model.CardEvent) error {
	query := psql.Insert("cards_events").
		Columns("card_id, type, status, payload, updated_at")

	updated := time.Now()
	for _, event := range events {
		if payload, err := json.Marshal(event.Entity); err == nil {
			query = query.Values(event.ID, event.Type, event.Status, payload, updated)
		} else {
			return err
		}
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	return err
}

func (r repo) Remove(eventIDs []uint64) error {
	query := sq.Delete("card_events").Where(sq.Eq{"id": eventIDs})
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(s, args...)
	return err
}
