
-- +migrate Up
CREATE TABLE IF NOT EXISTS chat_posts
(
    chat_post_id TEXT NOT NULL,
    room_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(chat_post_id)
);
-- +migrate Down
DROP TABLE IF EXISTS chat_posts;