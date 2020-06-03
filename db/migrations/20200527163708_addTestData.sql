
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('54ef66a2-c469-4410-8516-ccb500145a70', 'test1', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('64ef66a2-c469-4410-8516-ccb500145a70', 'test2', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('74ef66a2-c469-4410-8516-ccb500145a70', 'test3', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('84ef66a2-c469-4410-8516-ccb500145a70', 'test4', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('94ef66a2-c469-4410-8516-ccb500145a70', 'test5', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('14ef66a2-c469-4410-8516-ccb500145a70', 'test6', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('24ef66a2-c469-4410-8516-ccb500145a70', 'test7', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('34ef66a2-c469-4410-8516-ccb500145a70', 'test8', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

INSERT INTO "game_titles"
    ("game_title_id","game_title","created_at","update_at")
VALUES
    ('44ef66a2-c469-4410-8516-ccb500145a70', 'test9', '2020-05-27 16:38:04', '2020-05-27 16:38:04');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

