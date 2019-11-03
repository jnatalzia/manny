import { getJSONFromPath } from '../utils/fetch';

export let currentGamestate = {
    land: {},
    players: {},
    traders: {}
};

export function initGamestate() {
    return getJSONFromPath('/gamestate?gid=7&pid=0').then(gs => {
        gs.land_tiles.reduce((_, lt) => {
            currentGamestate.land[lt.id] = Object.assign(lt, {
                traders: []
            });
        }, {});

        gs.current_player.land_tiles.reduce((_, lt) => {
            let existingLand = currentGamestate.land[lt.id] || {};
            currentGamestate.land[lt.id] = Object.assign(existingLand, lt, {
                current_player_owns: true
            });
        }, {});

        gs.players.reduce((_, p) => {
            currentGamestate.players[p.id] = p;
        }, {});

        currentGamestate.players[gs.current_player.id] = Object.assign(
            gs.current_player,
            {
                current_player_owns: true
            }
        );

        gs.traders.reduce((_, t) => {
            currentGamestate.traders[t.id] = t;
            currentGamestate.traders[t.id].landArrIdx =
                currentGamestate.land[t.location_id].traders.length;
            currentGamestate.land[t.location_id].traders.push(
                currentGamestate.traders[t.id]
            );
        }, {});

        gs.current_player.traders.reduce((_, t) => {
            let existingLand = currentGamestate.traders[t.id] || {};
            currentGamestate.traders[t.id] = Object.assign(existingLand, t, {
                current_player_owns: true
            });

            currentGamestate.land[t.location_id].traders.splice(
                currentGamestate.traders[t.id].landArrIdx,
                1,
                currentGamestate.traders[t.id]
            );
        }, {});

        currentGamestate.turn_number = gs.turn_number;
        console.log(currentGamestate);
    });
}
