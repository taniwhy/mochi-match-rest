
-- +migrate Up
CREATE TABLE IF NOT EXISTS rooms
(
    room_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    room_text TEXT NOT NULL,
    game_list_id TEXT NOT NULL,
    game_hard_id TEXT NOT NULL,
    capacity INTEGER NOT NULL,
    is_lock BOOLEAN NOT NULL,
    start TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(room_id),
    FOREIGN KEY(user_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(game_list_id)REFERENCES game_lists(game_list_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(game_hard_id)REFERENCES game_hards(game_hard_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);
-- +migrate Down
DROP TABLE IF EXISTS rooms;