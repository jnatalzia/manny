-- Verify manipulation:table.initial_schemas on pg

BEGIN;

-- XXX Add verifications here.

SELECT * FROM games;

SELECT * FROM game_player_join;

SELECT * FROM players;

ROLLBACK;