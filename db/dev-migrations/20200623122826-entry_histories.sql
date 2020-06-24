
-- +migrate Up
CREATE TABLE IF NOT EXISTS entry_histories
(
    entry_history_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    room_id TEXT NOT NULL,
    is_leave BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    leaved_at TIMESTAMP,
    PRIMARY KEY(entry_history_id),
    FOREIGN KEY(user_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(room_id)REFERENCES rooms(room_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE IF EXISTS entry_histories;