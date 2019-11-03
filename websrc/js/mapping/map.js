import mapJSON from './map.json';
import { camera } from './camera.js';
import { CANVAS_WIDTH, CANVAS_HEIGHT } from '../components/ctx.js';

export function drawMap(ctx) {
    mapJSON.forEach(land => {
        ctx.save();
        ctx.translate(CANVAS_WIDTH / 2, CANVAS_HEIGHT / 2);
        ctx.translate(
            -camera.center.x * camera.zoomLevel,
            -camera.center.y * camera.zoomLevel
        );
        ctx.scale(camera.zoomLevel, camera.zoomLevel);

        ctx.beginPath();
        ctx.moveTo(land.coords[0].x, land.coords[0].y);

        land.coords.slice(1).forEach(c => {
            ctx.lineTo(c.x, c.y);
        });

        ctx.closePath();
        ctx.stroke();

        ctx.restore();
    });
}
