package api

import (
	"context"
	"github.com/ozonmp/pay-card-api/internal/model"
	"github.com/ozonmp/pay-card-api/internal/repo/cards"
	repo_cards_events "github.com/ozonmp/pay-card-api/internal/repo/cards_events"
	"github.com/rs/zerolog/log"
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

func (a cardAPI) CreateCard(ctx context.Context, req *pb.Card) (*pb.CreateCardV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateCardV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("CreateCard request happens")

	createEvent := MapCreateEvent(req)

	id, err := a.repo.Add(createEvent.Entity)
	if err != nil {
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
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveCardV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("RemoveCard request happens")

	removeEvent := MapRemoveEvent(req)

	_, err := a.repo.Remove(removeEvent.Entity.CardId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = a.repoEvents.Add([]model.CardEvent{*removeEvent})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, err
}

func (a cardAPI) DescribeCard(ctx context.Context, req *pb.DescribeCardV1Request) (*pb.Card, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeCardV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("DescribeCard request happens")

	card, err := a.repo.Get(req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListCard - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().Msg("ListCard request happens")

	cards, err := a.repo.List(req.GetOffset(), req.GetLimit())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := MapProtoListModel(cards)

	return response, err
}
