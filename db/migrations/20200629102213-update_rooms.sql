
-- +migrate Up
ALTER TABLE rooms ALTER COLUMN game_list_id DROP NOT NULL;
ALTER TABLE rooms ALTER COLUMN game_hard_id DROP NOT NULL;

-- +migrate Down
ALTER TABLE rooms ALTER COLUMN game_list_id SET NOT NULL;
ALTER TABLE rooms ALTER COLUMN game_hard_id SET NOT NULL;


INSERT INTO "entry_histories" 
	("entry_history_id","user_id","room_id","is_leave","created_at","leaved_at") 
VALUES (
    ca4ad7a6-e118-4acf-8692-76e8379de2f8
    4a48ce84-3a4b-4d27-bd25-ec0140c0a4f5 
    c644af10-690d-47fa-8b7b-752e0decd870 