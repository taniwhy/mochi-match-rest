
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
    id SERIAL,
    user_name VARCHAR,
    provider TEXT,
    provider_id TEXT,
    is_admin BOOLEAN,
    is_frozen BOOLEAN,
    created_at TIMESTAMPTZ,
    update_at TIMESTAMPTZ,
    PRIMARY KEY(id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users