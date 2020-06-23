
-- +migrate Up
CREATE TABLE IF NOT EXISTS game_hards
(
    game_hard_id TEXT NOT NULL,
    game_hard TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(game_hard_id),
    UNIQUE(game_hard)
);
-- +migrate Down
DROP TABLE IF EXISTS game_hards;