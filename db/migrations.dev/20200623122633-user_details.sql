
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_details
(
    user_detail_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    icon TEXT NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(user_detail_id),
    UNIQUE(user_id),
    FOREIGN KEY(user_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS user_details;