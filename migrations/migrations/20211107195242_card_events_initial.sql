-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cards
(
    id              SERIAL PRIMARY KEY,
    owner_id        BIGINT       NOT NULL,
    payment_system  VARCHAR(255) NOT NULL,
    number          VARCHAR(255) NOT NULL,
    holder_name     VARCHAR(255) NOT NULL,
    expiration_date TIMESTAMP    NOT NULL,
    CvcCvv          VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS cards_events
(
    id         SERIAL PRIMARY KEY,
    card_id    BIGINT       NOT NULL,
    type       VARCHAR(255) NOT NULL,
    status     VARCHAR(255) NOT NULL,
    payload    JSONB        NULL,
    updated_at TIMESTAMP    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS cards_events;
-- +goose StatementEnd
