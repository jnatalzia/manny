import { drawMap, getTileAtMousePoint } from './mapping/map';
import { updateCamera, moveCameraToFocus } from './mapping/camera';
import { didClickThisCycle, getClickPoint } from './components/controls';
import { openTileStats } from './components/modals/tile_stats';
import { currentGamestate } from './components/gamestate';

function checkClick() {
    if (didClickThisCycle()) {
        const pt = getClickPoint();
        moveCameraToFocus(pt);
        let lID = getTileAtMousePoint(pt);
        if (typeof lID !== 'undefined') {
            openTileStats(currentGamestate.land[lID]);
        }
    }
}

function updateGame(delta) {
    updateCamera();
    checkClick();
}

function drawGame(ctx) {
    ctx.clearRect(0, 0, window.innerWidth, window.innerHeight);
    drawMap(ctx);
}

export function update(ctx) {
    updateGame(ctx);
    drawGame(ctx);

    // redraw
    window.requestAnimationFrame(update.bind(null, ctx));
}
