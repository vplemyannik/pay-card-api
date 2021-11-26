package model

import (
	"time"
)

type Card struct {
	CardId         uint64    `db:"id"`
	OwnerId        uint64    `db:"owner_id"`
	PaymentSystem  string    `db:"payment_system"`
	Number         string    `db:"number"`
	HolderName     string    `db:"holder_name"`
	ExpirationDate time.Time `db:"expiration_date"`
	CvcCvv         string    `db:"cvccvv"`
}

type EventType string

type EventStatus string

const (
	Created EventType = "Created"
	Updated EventType = "Updated"
	Removed EventType = "Removed"
)

const (
	New       EventStatus = "New"
	Locked    EventStatus = "Locked"
	Processed EventStatus = "Processed"
)

type CardEventPayload interface {
	GetCardId() uint64
}

type CreateCardEventPayload struct {
	CardId         uint64
	OwnerId        uint64
	PaymentSystem  string
	Number         string
	HolderName     string
	ExpirationDate time.Time
	CvcCvv         string
}

type UpdateCardEventPayload struct {
	CardId         uint64
	OwnerId        *uint64
	PaymentSystem  *string
	Number         *string
	HolderName     *string
	ExpirationDate *time.Time
	CvcCvv         *string
}

type RemoveCardEventPayload struct {
	CardId uint64
}

type CardEvent struct {
	ID        uint64
	Type      EventType
	Status    EventStatus
	Entity    interface{}
	OccuredAt time.Time
}

func (c CreateCardEventPayload) MapToCard() Card {
	return Card{
		CardId:         c.CardId,
		OwnerId:        c.OwnerId,
		PaymentSystem:  c.PaymentSystem,
		Number:         c.Number,
		HolderName:     c.HolderName,
		ExpirationDate: c.ExpirationDate,
		CvcCvv:         c.CvcCvv,
	}
}

func (c CreateCardEventPayload) GetCardId() uint64 {
	return c.CardId
}

func (c UpdateCardEventPayload) GetCardId() uint64 {
	return c.CardId
}

func (c RemoveCardEventPayload) GetCardId() uint64 {
	return c.CardId
}
