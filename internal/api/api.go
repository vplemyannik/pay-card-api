package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/pay-card-api/internal/model"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
	"github.com/ozonmp/pay-card-api/internal/pkg/metrics"
	"github.com/ozonmp/pay-card-api/internal/repo/cards"
	repo_cards_events "github.com/ozonmp/pay-card-api/internal/repo/cards_events"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	pb "github.com/ozonmp/pay-card-api/pkg/pay-card-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "omp_template_api_template_not_found_total",
		Help: "Total number of templates that were not found",
	})
)

type cardAPI struct {
	pb.UnimplementedPayCardApiServiceServer
	repo       repo_cards.Repo
	repoEvents repo_cards_events.Repo
}

func NewTemplateAPI(r repo_cards.Repo, repoEvents repo_cards_events.Repo) pb.PayCardApiServiceServer {
	return &cardAPI{repo: r, repoEvents: repoEvents}
}

func (a cardAPI) CreateCard(ctx context.Context, req *pb.CreateCardV1Request) (*pb.CreateCardV1Response, error) {
	notifySpan, _ := opentracing.StartSpanFromContext(ctx, "CreateCardV1Request - start request")
	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "CreateCardV1Request - invalid argument", "err", err)
		notifySpan.SetTag("CreateCardV1Request - invalid argument", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "CreateCard request happens")
	metrics.IncrementCUDCardOperationsTotalCount(metrics.Create)

	createEventPayload := MapCreateCardEventPayload(req)
	createEvent := model.CardEvent{
		Type:      model.Created,
		Status:    model.New,
		Entity:    createEventPayload,
		OccuredAt: time.Now(),
	}

	card := createEventPayload.MapToCard()
	id, err := a.repo.Add(ctx, &card)
	createEventPayload.CardId = id

	if err != nil {
		card := req.GetCard()
		notifySpan.SetTag("CreateCardV1Request - error occured when create card", err).
			SetTag("HolderName", card.GetHolderName()).
			SetTag("Cvc", card.GetCvcCvv()).
			SetTag("Number", card.GetNumber()).
			SetTag("OwnerId", card.GetOwnerId()).
			SetTag("PaymentSystem", card.GetPaymentSystem())
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = a.repoEvents.Add([]model.CardEvent{createEvent})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.CreateCardV1Response{
		Id: id,
	}

	return &response, err
}

func (a cardAPI) UpdateCard(ctx context.Context, req *pb.UpdateCardV1Request) (*emptypb.Empty, error) {
	notifySpan, _ := opentracing.StartSpanFromContext(ctx, "UpdateCardV1Request - start request")
	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "UpdateCardV1Request - invalid argument", "err", err)
		notifySpan.SetTag("UpdateCardV1Request - invalid argument", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "UpdateCard request happens")
	metrics.IncrementCUDCardOperationsTotalCount(metrics.Update)

	updateEventPayload := MapUpdateCardEventPayload(req)
	updateEvent := model.CardEvent{
		Type:      model.Updated,
		Status:    model.New,
		Entity:    updateEventPayload,
		OccuredAt: time.Now(),
	}

	card, err := a.repo.Get(updateEventPayload.GetCardId())
	if updateEventPayload.OwnerId != nil {
		card.OwnerId = *updateEventPayload.OwnerId
	}
	if updateEventPayload.PaymentSystem != nil {
		card.PaymentSystem = *updateEventPayload.PaymentSystem
	}
	if updateEventPayload.Number != nil {
		card.Number = *updateEventPayload.Number
	}
	if updateEventPayload.HolderName != nil {
		card.HolderName = *updateEventPayload.HolderName
	}
	if updateEventPayload.ExpirationDate != nil {
		card.ExpirationDate = *updateEventPayload.ExpirationDate
	}
	if updateEventPayload.CvcCvv != nil {
		card.CvcCvv = *updateEventPayload.CvcCvv
	}

	err = a.repo.Update(ctx, card)

	if err != nil {
		card := updateEventPayload
		notifySpan.SetTag("UpdateCardV1Request - error occured when update card", err)
		if card.OwnerId != nil {
			notifySpan = notifySpan.SetTag("owner_id", *card.OwnerId)
		}
		if card.PaymentSystem != nil {
			notifySpan = notifySpan.SetTag("payment_system", *card.PaymentSystem)
		}
		if card.Number != nil {
			notifySpan = notifySpan.SetTag("number", *card.Number)
		}
		if card.HolderName != nil {
			notifySpan = notifySpan.SetTag("holder_name", *card.HolderName)
		}
		if card.ExpirationDate != nil {
			notifySpan = notifySpan.SetTag("expiration_date", *card.ExpirationDate)
		}
		if card.CvcCvv != nil {
			notifySpan = notifySpan.SetTag("cvccvv", *card.CvcCvv)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = a.repoEvents.Add([]model.CardEvent{updateEvent})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, err
}

func (a cardAPI) RemoveCard(ctx context.Context, req *pb.RemoveCardV1Request) (*emptypb.Empty, error) {
	notifySpan, _ := opentracing.StartSpanFromContext(ctx, "RemoveCard - start request")
	defer notifySpan.Finish()

	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "RemoveCardV1Request - invalid argument", "err", err)
		notifySpan.SetTag("RemoveCardV1Request - invalid argument", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "RemoveCard request happens")
	metrics.IncrementCUDCardOperationsTotalCount(metrics.Delete)

	removeEventPayload := MapRemoveCardEventPayload(req)
	removeEvent := model.CardEvent{
		Type:      model.Removed,
		Status:    model.New,
		Entity:    removeEventPayload,
		OccuredAt: time.Now(),
	}

	_, err := a.repo.Remove(removeEventPayload.CardId)
	if err != nil {
		notifySpan.SetTag("RemoveCardV1Request - remove error", err).
			SetTag("CardId", req.GetId())
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = a.repoEvents.Add([]model.CardEvent{removeEvent})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, err
}

func (a cardAPI) DescribeCard(ctx context.Context, req *pb.DescribeCardV1Request) (*pb.Card, error) {
	notifySpan, _ := opentracing.StartSpanFromContext(ctx, "DescribeCard - start request")
	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "DescribeCardV1Request - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "DescribeCard request happens")

	card, err := a.repo.Get(req.GetId())
	if err != nil {
		notifySpan.SetTag("DescribeCardV1Request - get card by id error", err).
			SetTag("CardId", req.GetId())
		return nil, status.Error(codes.Internal, err.Error())
	}

	if card == nil {
		metrics.IncrementNotFoundCardTotalCount()
		return nil, status.Error(codes.NotFound, "card not found")
	}

	response := &pb.Card{
		OwnerId:        card.OwnerId,
		PaymentSystem:  card.PaymentSystem,
		Number:         card.Number,
		HolderName:     card.HolderName,
		CvcCvv:         card.CvcCvv,
		ExpirationDate: timestamppb.New(card.ExpirationDate),
	}

	return response, err
}

func (a cardAPI) ListCard(ctx context.Context, req *pb.ListCardV1Request) (*pb.ListCardV1Response, error) {
	notifySpan, _ := opentracing.StartSpanFromContext(ctx, "ListCard - start request")
	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "ListCardV1Request - invalid argument", "err", err)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logger.InfoKV(ctx, "ListCard request happens")

	cards, err := a.repo.List(req.GetOffset(), req.GetLimit())
	if err != nil {
		notifySpan.SetTag("ListCardV1Request - get cards error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := MapProtoListModel(cards)

	return response, err
}
