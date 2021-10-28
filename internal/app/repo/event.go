package repo

import "github.com/ozonmp/pay-card-api/internal/model"

type EventRepo interface {
	Lock(n uint64) ([]model.CardEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.CardEvent) error
	Remove(eventIDs []uint64) error
}
