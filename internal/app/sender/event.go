package sender

import (
	"github.com/Shopify/sarama"
	"github.com/ozonmp/pay-card-api/internal/model"
	events "github.com/ozonmp/pay-card-api/pkg/pay-card-events"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type EventSender interface {
	Send(subdomain []model.CardEvent) error
}

type kafkaSender struct {
	producer sarama.SyncProducer
}

func NewKafkaSender(producer sarama.SyncProducer) EventSender {
	return &kafkaSender{
		producer: producer,
	}
}

func (sender kafkaSender) Send(events []model.CardEvent) error {
	for _, event := range events {
		body, _ := marshalEvent(event)
		switch event.Type {
		case model.Created:
			SendMessage(sender.producer, "card_created", body)
		case model.Removed:
			SendMessage(sender.producer, "card_removed", body)
		case model.Updated:
			SendMessage(sender.producer, "card_updated", body)
		}
	}

	return nil
}

func marshalEvent(event model.CardEvent) ([]byte, error) {
	var protoMessage proto.Message
	switch event.Type {
	case model.Created:
		protoMessage = mapIntoCreatedProto(event.Entity, event.OccuredAt)
	case model.Removed:
		protoMessage = mapIntoRemovedProto(event.Entity.CardId, event.OccuredAt)
	case model.Updated:
		protoMessage = mapIntoUpdatedProto(event.Entity, event.OccuredAt)
	}
	return protojson.Marshal(protoMessage)
}

func mapIntoCreatedProto(card model.Card, createdAt time.Time) *events.CardCreated {
	return &events.CardCreated{
		CreatedAt: timestamppb.New(createdAt),
		Card: &events.Card{
			OwnerId:        card.OwnerId,
			PaymentSystem:  card.PaymentSystem,
			Number:         card.Number,
			HolderName:     card.HolderName,
			CvcCvv:         card.CvcCvv,
			CardId:         card.CardId,
			ExpirationDate: timestamppb.New(card.ExpirationDate),
		},
	}
}

func mapIntoRemovedProto(cardId uint64, createdAt time.Time) *events.CardDeleted {
	return &events.CardDeleted{
		RemovedAt: timestamppb.New(createdAt),
		CardId:    cardId,
	}
}

func mapIntoUpdatedProto(card model.Card, updatedAt time.Time) *events.CardUpdated {

	return &events.CardUpdated{
		UpdatedAt: timestamppb.New(updatedAt),
		Card: &events.UpdateCard{
			OwnerId:        &card.OwnerId,
			PaymentSystem:  card.PaymentSystem,
			Number:         card.Number,
			HolderName:     card.HolderName,
			CvcCvv:         card.CvcCvv,
			CardId:         card.CardId,
			ExpirationDate: timestamppb.New(card.ExpirationDate),
		},
	}
}
