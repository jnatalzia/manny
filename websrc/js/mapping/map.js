import mapJSON from './map.json';
import { camera } from './camera.js';
import { CANVAS_WIDTH, CANVAS_HEIGHT } from '../components/ctx.js';
import { currentGamestate } from '../components/gamestate.js';
import { doesOverlap } from '../utils/intersection.js';

export const MAP_WIDTH = 500;

function getLandColorForID(id) {
    let owner = currentGamestate.players[currentGamestate.land[id].owner_id];
    return owner ? owner.land_color : '#ddd';
}

function drawMetadata(land, ctx) {
    let landData = currentGamestate.land[land.landID];
    let numTraders = landData.traders.length;
    if (numTraders === 0) {
        return;
    }
    const startingCoords = land.coords[0];
    ctx.save();
    ctx.font = '14pt Arial';
    ctx.textBaseline = 'middle';
    ctx.textAlign = 'center';
    ctx.fillStyle = '#111';
    ctx.fillText(
        `${numTraders}T`,
        startingCoords.x + 50,
        startingCoords.y + 50
    );
    ctx.restore();
}

export function getTileAtMousePoint(pt) {
    // let adjustedPT =
    let zCenterX = camera.center.x;
    let zCenterY = camera.center.y;

    let box = {
        x: zCenterX,
        y: zCenterY,
        size: 1
    };

    let overlap;
    for (let i = 0; i < mapJSON.length; i++) {
        if (doesOverlap(box, mapJSON[i])) {
            overlap = mapJSON[i];
        }
    }
    if (!overlap) {
        return;
    }
    return overlap.landID;
}

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
        ctx.fillStyle = getLandColorForID(land.landID);
        ctx.strokeStyle = '#aaa';
        ctx.fill();
        ctx.stroke();

        drawMetadata(land, ctx);

        ctx.restore();
    });
}
