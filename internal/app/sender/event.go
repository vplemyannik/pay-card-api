package sender

import "github.com/ozonmp/pay-card-api/internal/model"

type EventSender interface {
	Send(subdomain *model.CardEvent) error
}
