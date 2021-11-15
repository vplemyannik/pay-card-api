package api

import (
	"github.com/ozonmp/pay-card-api/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ozonmp/pay-card-api/pkg/pay-card-api"
)

func MapCreateEvent(req *pb.CreateCardV1Request) *model.CardEvent {
	card := req.GetCard()
	entity := &model.Card{
		OwnerId:        card.GetOwnerId(),
		PaymentSystem:  card.GetPaymentSystem(),
		ExpirationDate: card.GetExpirationDate().AsTime(),
		HolderName:     card.GetHolderName(),
		CvcCvv:         card.GetCvcCvv(),
		Number:         card.GetNumber(),
	}

	return &model.CardEvent{
		Type:   model.Created,
		Status: model.New,
		Entity: entity,
	}
}

func MapRemoveEvent(r *pb.RemoveCardV1Request) *model.CardEvent {
	cardId := r.GetId()
	return &model.CardEvent{
		Type:   model.Removed,
		Status: model.New,
		Entity: &model.Card{
			CardId: cardId,
		},
	}
}

func MapProtoModel(card *model.Card) *pb.Card {
	return &pb.Card{
		OwnerId:        card.OwnerId,
		PaymentSystem:  card.PaymentSystem,
		Number:         card.Number,
		HolderName:     card.HolderName,
		CvcCvv:         card.CvcCvv,
		ExpirationDate: timestamppb.New(card.ExpirationDate),
	}
}

func MapProtoListModel(cardModels []model.Card) *pb.ListCardV1Response {
	var cards = make([]*pb.Card, 0, len(cardModels))

	for _, card := range cardModels {
		cards = append(cards, MapProtoModel(&card))
	}

	return &pb.ListCardV1Response{
		Cards: cards,
	}
}
