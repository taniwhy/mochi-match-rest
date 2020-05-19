
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
    user_id BIGSERIAL NOT NULL,
    google_id TEXT,
    twitter_id TEXT,
    facebook_id TEXT,
    is_admin BOOLEAN NOT NULL,
    is_freeze BOOLEAN NOT NULL,
    is_delete BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    update_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZNOT NULL,
    PRIMARY KEY(user_id),
    UNIQUE(google_id,twitter_id,facebook_id)
);

CREATE TABLE user_detail
(
    user_detail_id BIGSERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    user_name TEXT NOT NULL,
    icon INTEGER NOT NULL,
    update_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(user_detail_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE favorate_game
(
    favorate_game_id BIGSERIAL NOT NULL,
    user_detail_id INTEGER NOT NULL,
    game_title INTEGER NOT NULL,
    update_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(user_detail_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY(game_title) REFERENCES game_title(game_title_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE game_title
(
    game_title_id BIGSERIAL NOT NULL,
    game_title TEXT NOT NULL,
    update_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(user_detail_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE room
(
    room_id BIGSERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    game_title_id INTEGER NOT NULL,
    capacity INTEGER NOT NULL,
    lock_flg BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(room_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY(game_title_id) REFERENCES game_title(game_title_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE game_title;
DROP TABLE favorate_game;
DROP TABLE user_detail;
DROP TABLE users;



