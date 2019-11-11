import '../css/main.scss';
import { ctx } from './components/ctx';
import { update } from './gameloop';
import { initControls } from './components/controls';
import { initGamestate, currentGamestate } from './components/gamestate';
import { openTileStats } from './components/modals/tile_stats';
import { addGenericModalListeners } from './components/modals/generic';

window.addEventListener('load', () => {
    initControls();
    addGenericModalListeners();
    initGamestate().then(() => {
        update(ctx);
    });
});
