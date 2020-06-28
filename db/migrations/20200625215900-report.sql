
-- +migrate Up
CREATE TABLE IF NOT EXISTS reports
(
    report_id TEXT NOT NULL,
    caller_id TEXT NOT NULL,
    vaiolator_id TEXT NOT NULL,
    reported_room_id TEXT NOT NULL,
    vaiolation_detail TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY(report_id),
    FOREIGN KEY(caller_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(vaiolator_id)REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY(reported_room_id)REFERENCES rooms(room_id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);

-- +migrate Down
DROP TABLE IF EXISTS reports;