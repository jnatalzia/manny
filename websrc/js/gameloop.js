import { drawMap } from './mapping/map';
import { updateCamera } from './mapping/camera';

function updateGame(delta) {
    updateCamera();
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
