-- CREATE DATABASE video;
\c video

CREATE TABLE videos (
    "id" BIGSERIAL NOT NULL,
    "user_id" VARCHAR(50),
    "title" VARCHAR(50),
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE rates (
    "id" BIGSERIAL UNIQUE NOT NULL,
    "user_id" VARCHAR(50),
    "video_id" BIGINT,
    "value" DECIMAL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("user_id", "video_id")
);

CREATE TABLE analyses (
    "id" BIGSERIAL UNIQUE NOT NULL,
    "user_id" VARCHAR(50),
    "video_id" BIGINT,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    PRIMARY KEY ("id")
);

-- テストデータ
INSERT INTO videos VALUES(1, 'user_1', 'title 1', '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO videos VALUES(2, 'user_2', 'title 2', '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO videos VALUES(3, 'user_3', 'title 3', '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO videos VALUES(4, 'user_4', 'title 4', '2000-01-01 00:00:00', '2000-01-01 00:00:00');

INSERT INTO rates VALUES(1, 'user_1', 1, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO rates VALUES(2, 'user_2', 2, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO rates VALUES(3, 'user_3', 3, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO rates VALUES(4, 'user_4', 4, 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');

INSERT INTO analyses VALUES(1, 'user_1', 1, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO analyses VALUES(2, 'user_2', 2, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO analyses VALUES(3, 'user_3', 3, '2000-01-01 00:00:00', '2000-01-01 00:00:00');
INSERT INTO analyses VALUES(4, 'user_4', 4, '2000-01-01 00:00:00', '2000-01-01 00:00:00');

SELECT setval('videos_id_seq', (SELECT MAX(id) FROM videos));
SELECT setval('rates_id_seq', (SELECT MAX(id) FROM rates));
SELECT setval('analyses_id_seq', (SELECT MAX(id) FROM analyses));
