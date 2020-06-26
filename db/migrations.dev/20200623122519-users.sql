
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    user_id TEXT NOT NULL,
    google_id TEXT,
    facebook_id TEXT,
    twitter_id TEXT,
    email TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL,
    is_freeze BOOLEAN NOT NULL,
    is_delete BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    delete_at TIMESTAMP,
    PRIMARY KEY(user_id),
    UNIQUE(google_id, facebook_id, twitter_id, email)
);
-- +migrate Down
DROP TABLE IF EXISTS users;