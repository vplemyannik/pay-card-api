package repo_cards_events

import (
	"database/sql"
	"github.com/ozonmp/pay-card-api/internal/model"
	"time"
)

type CardEventDb struct {
	ID        uint64            `db:"id"`
	CardId    uint64            `db:"card_id"`
	Type      model.EventType   `db:"type"`
	Status    model.EventStatus `db:"status"`
	Payload   sql.NullString    `db:"payload"`
	UpdatedAt time.Time         `db:"updated_at"`
}
