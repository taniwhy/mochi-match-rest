
-- +migrate Up
CREATE TABLE IF NOT EXISTS game_hards
(
    game_hard_id TEXT NOT NULL,
    hard_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(game_hard_id),
    UNIQUE(hard_name)
);
-- +migrate Down
DROP TABLE IF EXISTS game_hards;