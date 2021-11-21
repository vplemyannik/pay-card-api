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

type CardEvent struct {
	ID        uint64
	Type      EventType
	Status    EventStatus
	Entity    Card
	OccuredAt time.Time
}
