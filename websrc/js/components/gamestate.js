import { getJSONFromPath } from '../utils/fetch';

export let massagedGameState = {
    land: {},
    players: {},
    current_player: {},
    traders: {}
};

export function initGamestate() {
    return getJSONFromPath('/gamestate?gid=6&pid=0').then(gs => {
        console.log(gs);
        gs.land_tiles.reduce((_, lt) => {
            massagedGameState.land[lt.id] = lt;
        }, {});

        gs.current_player.land_tiles.reduce((_, lt) => {
            let existingLand = massagedGameState.land[lt.id] || {};
            massagedGameState.land[lt.id] = Object.assign(existingLand, lt, {
                current_player_owns: true
            });
        }, {});

        massagedGameState.players = gs.players;

        gs.traders.reduce((_, t) => {
            massagedGameState.traders[t.id] = t;
        }, {});

        gs.current_player.traders.reduce((_, t) => {
            let existingLand = massagedGameState.traders[t.id] || {};
            massagedGameState.traders[t.id] = Object.assign(existingLand, t, {
                current_player_owns: true
            });
        }, {});

        massagedGameState.turn_number = gs.turn_number;
        massagedGameState.current_player = gs.current_player;
        delete massagedGameState.current_player.land_tiles;
        delete massagedGameState.current_player.traders;
        console.log(massagedGameState);
    });
}
