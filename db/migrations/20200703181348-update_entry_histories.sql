
-- +migrate Up
ALTER TABLE entry_histories ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE entry_histories ALTER COLUMN room_id DROP NOT NULL;

-- +migrate Down
ALTER TABLE entry_histories ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE entry_histories ALTER COLUMN room_id SET NOT NULL;
