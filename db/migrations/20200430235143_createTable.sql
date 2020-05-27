
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users
(
    user_id TEXT NOT NULL,
    google_id TEXT,
    facebook_id TEXT,
    twitter_id TEXT,
    is_admin BOOLEAN NOT NULL,
    is_freeze BOOLEAN NOT NULL,
    is_delete BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    delete_at TIMESTAMP,
    PRIMARY KEY(user_id),
    UNIQUE(google_id, facebook_id, twitter_id)
);

CREATE TABLE user_details
(
    user_detail_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    user_name TEXT NOT NULL,
    icon INTEGER NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(user_detail_id),
    FOREIGN KEY(user_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE chat_posts
(
    chat_post_id TEXT NOT NULL,
    room_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(chat_post_id)
);

CREATE TABLE game_titles
(
    game_title_id TEXT NOT NULL,
    game_title TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL,
    PRIMARY KEY(game_title_id)
);

CREATE TABLE rooms
(
    room_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    game_title_id TEXT NOT NULL,
    capacity INTEGER NOT NULL,
    PRIMARY KEY(room_id),
    FOREIGN KEY(user_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(game_title_id)REFERENCES game_titles(game_title_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS chat_posts;
DROP TABLE IF EXISTS favorate_game;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS user_details;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS game_titles;





