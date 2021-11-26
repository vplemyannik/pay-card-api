package repo_cards_events

import (
	"context"
	"database/sql"
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

	events := make([]model.CardEvent, 0, n)

	ctx := context.Background()
	err := db.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) error {
		isAcquired, err := db.AcquireTryLock(ctx, tx, int32(lockID), int32(entityID))
		if !isAcquired || err != nil {
			return errors.Wrap(err, "database.WithTx")
		}

		query := psql.Update("cards_events").
			Set("status", model.Locked).
			Where(sq.Select("id").
				Prefix("id IN (").
				From("cards_events").
				Where(sq.Or{
					sq.Eq{"status": model.New},
				}).
				OrderBy("updated_at ASC").
				Limit(n).
				Suffix(")")).
			Suffix("RETURNING *").
			RunWith(r.db)

		sql, args, err := query.ToSql()
		if err != nil {
			return err
		}

		var res []CardEventDb
		err = r.db.Select(&res, sql, args...)
		if err != nil {
			return err
		}

		for _, event := range res {
			mappedEvent, err := mapToDomain(event)
			if err != nil {
				return err
			}
			events = append(events, *mappedEvent)
		}

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
			query = query.Values(event.Entity.(CardEventPayload).GetCardId(), event.Type, event.Status, payload, updated)
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
	query := psql.Delete("cards_events").
		Where(sq.Eq{"id": eventIDs}).
		RunWith(r.db)

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(s, args...)
	return err
}

func mapToDomain(eventDb CardEventDb) (*model.CardEvent, error) {
	entity, err := UnmarshalEntity(eventDb.Type, eventDb.Payload)
	if err != nil {
		return nil, err
	}

	return &CardEvent{
		ID:        eventDb.ID,
		Type:      eventDb.Type,
		Status:    eventDb.Status,
		Entity:    entity,
		OccuredAt: eventDb.UpdatedAt,
	}, nil
}

func UnmarshalEntity(eventType model.EventType, payload sql.NullString) (interface{}, error) {
	switch eventType {
	case Created:
		var created model.CreateCardEventPayload
		err := UnmarshalEvent(payload, &created)
		return created, err
	case Removed:
		var removed model.RemoveCardEventPayload
		err := UnmarshalEvent(payload, &removed)
		return removed, err
	case Updated:
		var updated model.UpdateCardEventPayload
		err := UnmarshalEvent(payload, &updated)
		return updated, err
	}
	return nil, nil
}

func UnmarshalEvent(payload sql.NullString, value interface{}) error {
	if payload.Valid {
		if err := json.Unmarshal([]byte(payload.String), value); err != nil {
			return err
		}
	}
	return nil
}
