package api

import (
	"github.com/ozonmp/pay-card-api/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ozonmp/pay-card-api/pkg/pay-card-api"
)

func MapCreateCardEventPayload(req *pb.CreateCardV1Request) *model.CreateCardEventPayload {
	card := req.GetCard()
	return &model.CreateCardEventPayload{
		OwnerId:        card.GetOwnerId(),
		PaymentSystem:  card.GetPaymentSystem(),
		ExpirationDate: card.GetExpirationDate().AsTime(),
		HolderName:     card.GetHolderName(),
		CvcCvv:         card.GetCvcCvv(),
		Number:         card.GetNumber(),
	}
}

func MapRemoveCardEventPayload(r *pb.RemoveCardV1Request) *model.RemoveCardEventPayload {
	cardId := r.GetId()
	return &model.RemoveCardEventPayload{
		CardId: cardId,
	}
}

func MapUpdateCardEventPayload(r *pb.UpdateCardV1Request) *model.UpdateCardEventPayload {
	cardId := r.GetId()
	card := r.GetCard()

	var updateCard model.UpdateCardEventPayload
	updateCard.CardId = cardId
	if card.GetOwnerId() != nil {
		value := card.GetOwnerId().GetValue()
		updateCard.OwnerId = &value
	}
	if card.GetPaymentSystem() != nil {
		value := card.GetPaymentSystem().GetValue()
		updateCard.PaymentSystem = &value
	}
	if card.GetNumber() != nil {
		value := card.GetNumber().GetValue()
		updateCard.Number = &value
	}
	if card.GetHolderName() != nil {
		value := card.GetHolderName().GetValue()
		updateCard.HolderName = &value
	}
	if card.GetCvcCvv() != nil {
		value := card.GetCvcCvv().GetValue()
		updateCard.CvcCvv = &value
	}
	if card.GetExpirationDate() != nil {
		value := card.GetExpirationDate().AsTime()
		updateCard.ExpirationDate = &value
	}
	return &updateCard
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
