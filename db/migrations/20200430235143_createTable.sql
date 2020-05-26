
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE chat_posts
(
    chat_post_id TEXT NOT NULL,
    room INTEGER,
    user_id INTEGER,
    message TEXT,
    created_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS game_title;
DROP TABLE IF EXISTS favorate_game;
DROP TABLE IF EXISTS user_detail;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS chat_posts;



