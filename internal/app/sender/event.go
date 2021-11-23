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

func NewKafkaSender(brokers []string) (EventSender, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	return &kafkaSender{
		producer: producer,
	}, err
}

func (sender kafkaSender) Send(events []model.CardEvent) error {
	for _, event := range events {
		message, _ := createMessage(event)
		sender.producer.SendMessage(message)
	}
	return nil
}

func createMessage(event model.CardEvent) (*sarama.ProducerMessage, error) {
	var protoMessage proto.Message
	var topicName string
	switch event.Type {
	case model.Created:
		protoMessage = mapIntoCreatedProto(event.Entity.(model.CreateCardEventPayload), event.OccuredAt)
		topicName = "card_created"
	case model.Removed:
		protoMessage = mapIntoRemovedProto(event.Entity.(model.RemoveCardEventPayload).CardId, event.OccuredAt)
		topicName = "card_removed"
	case model.Updated:
		protoMessage = mapIntoUpdatedProto(event.Entity.(model.UpdateCardEventPayload), event.OccuredAt)
		topicName = "card_updated"
	}
	body, err := protojson.Marshal(protoMessage)
	return &sarama.ProducerMessage{
		Topic:     topicName,
		Partition: -1,
		Value:     sarama.ByteEncoder(body),
	}, err
}

func mapIntoCreatedProto(card model.CreateCardEventPayload, createdAt time.Time) *events.CardCreated {
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

func mapIntoUpdatedProto(card model.UpdateCardEventPayload, updatedAt time.Time) *events.CardUpdated {
	var expirationDate *timestamppb.Timestamp
	if card.ExpirationDate != nil {
		expirationDate = timestamppb.New(*card.ExpirationDate)
	}
	return &events.CardUpdated{
		UpdatedAt: timestamppb.New(updatedAt),
		Card: &events.UpdateCard{
			CardId:         card.CardId,
			OwnerId:        card.OwnerId,
			PaymentSystem:  card.PaymentSystem,
			Number:         card.Number,
			HolderName:     card.HolderName,
			CvcCvv:         card.CvcCvv,
			ExpirationDate: expirationDate,
		},
	}
}
