import { drawMap } from './mapping/map';
import { updateCamera } from './mapping/camera';

function updateGame(delta) {
    updateCamera();
}

function drawGame(ctx) {
    ctx.clearRect(0, 0, 1000, 1000);
    drawMap(ctx);
}

export function update(ctx) {
    updateGame(ctx);
    drawGame(ctx);

    window.requestAnimationFrame(update.bind(null, ctx));
}
