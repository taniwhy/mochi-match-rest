
-- +migrate Up
ALTER TABLE rooms ALTER COLUMN game_list_id DROP NOT NULL;
ALTER TABLE rooms ALTER COLUMN game_hard_id DROP NOT NULL;

-- +migrate Down
ALTER TABLE rooms ALTER COLUMN game_list_id SET NOT NULL;
ALTER TABLE rooms ALTER COLUMN game_hard_id SET NOT NULL;