package model

import "time"

type Card struct {
	OwnerId        uint64
	PaymentSystem  string
	Number         string
	HolderName     string
	ExpirationDate time.Time
	CvcCvv         string
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type CardEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Card
}
