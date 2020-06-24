
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
