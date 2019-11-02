-- Deploy manipulation:table.initial_schemas to pg

BEGIN;

-- XXX Add DDLs here.
CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    game_complete boolean,
    game_state json,
    pending_moves json
);

CREATE TABLE game_player_join (
    game_id INT,
    player_id INT,
    farm_id INT
);

CREATE TABLE players (
    player_id SERIAL PRIMARY KEY,
    username TEXT
);

COMMIT;
