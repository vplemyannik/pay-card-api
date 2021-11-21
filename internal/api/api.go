package api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/pay-card-api/internal/model"
	"github.com/ozonmp/pay-card-api/internal/pkg/logger"
	"github.com/ozonmp/pay-card-api/internal/pkg/metrics"
	"github.com/ozonmp/pay-card-api/internal/repo/cards"
	repo_cards_events "github.com/ozonmp/pay-card-api/internal/repo/cards_events"

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

	createEvent := MapCreateEvent(req)

	id, err := a.repo.Add(ctx, createEvent.Entity)
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
	err = a.repoEvents.Add([]model.CardEvent{*createEvent})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.CreateCardV1Response{
		Id: id,
	}

	return &response, err
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

	removeEvent := MapRemoveEvent(req)

	_, err := a.repo.Remove(removeEvent.Entity.CardId)
	if err != nil {
		notifySpan.SetTag("RemoveCardV1Request - remove error", err).
			SetTag("CardId", req.GetId())
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = a.repoEvents.Add([]model.CardEvent{*removeEvent})

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
