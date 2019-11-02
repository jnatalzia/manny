-- Revert manipulation:table.initial_schemas from pg

BEGIN;

-- XXX Add DDLs here.
DROP TABLE games;

DROP TABLE game_player_join;

DROP TABLE players;

COMMIT;
