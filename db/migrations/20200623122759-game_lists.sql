
-- +migrate Up
CREATE TABLE IF NOT EXISTS game_lists
(
    game_list_id TEXT NOT NULL,
    game_title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(game_list_id),
    UNIQUE(game_title)
);
-- +migrate Down
DROP TABLE IF EXISTS game_lists;