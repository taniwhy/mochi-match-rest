
-- +migrate Up

-- admin user
INSERT INTO users VALUES
(
    '1',
    '1',
    null,
    null,
    'tani@gmail.com',
    true,
    false,
    false,
    '2004-10-19 10:23:54',
    '2004-10-19 10:23:54',
    null
);

INSERT INTO user_details VALUES
(
    '1',
    '1',
    'taniwhy',
    'init',
    '2004-10-19 10:23:54'
);

-- common user
INSERT INTO users VALUES
(
    '2',
    '2',
    null,
    null,
    'tani2@gmail.com',
    false,
    false,
    false,
    '2004-10-19 10:23:54',
    '2004-10-19 10:23:54',
    null
);

INSERT INTO user_details VALUES
(
    '2',
    '2',
    'taniwhy',
    'init',
    '2004-10-19 10:23:54'
);

-- +migrate Down
DELETE FROM users WHERE user_id = '1';
DELETE FROM user_details WHERE user_id = '1';
DELETE FROM users WHERE user_id = '2';
DELETE FROM user_details WHERE user_id = '2';


SELECT
(
    rooms.room_id,
    rooms.user_id,
    user_details.icon,
    game_hards.hard_name,
    game_lists.game_title,
    rooms.capacity,
    rooms.room_text,
    user_details.user_name,
    (
        SELECT
        COUNT(entry_histories.entry_history_id)
        FROM entry_histories
        WHERE rooms.room_id = entry_histories.room_id
    ) As count,
    rooms.created_at,
    rooms.start
)
FROM "rooms"
LEFT JOIN user_details ON rooms.user_id = user_details.user_id
LEFT JOIN game_hards ON rooms.game_hard_id = game_hards.game_hard_id
LEFT JOIN game_lists ON rooms.game_list_id = game_lists.game_list_id
WHERE (rooms.is_lock = false)
ORDER BY created_at desc
LIMIT 8
OFFSET 0