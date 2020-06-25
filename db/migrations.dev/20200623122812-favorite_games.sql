
-- +migrate Up
CREATE TABLE IF NOT EXISTS favorite_games
(
    favorite_game_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    game_title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(favorite_game_id),
    FOREIGN KEY(user_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(game_title)REFERENCES game_lists(game_title)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE IF EXISTS favorite_games;