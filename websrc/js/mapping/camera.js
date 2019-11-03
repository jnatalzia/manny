import { KEY_MAP } from '../components/controls';
import { CANVAS_WIDTH, CANVAS_HEIGHT } from '../components/ctx';

const CAMERA_PAN_SPEED = 3;

export const camera = {
    zoomLevel: 1,
    center: {
        x: 250,
        y: 250
    }
};

export function updateCamera() {
    // LEFT
    if (KEY_MAP[37]) {
        camera.center.x -= 1 * CAMERA_PAN_SPEED;
    }
    // RIGHT
    if (KEY_MAP[39]) {
        camera.center.x += 1 * CAMERA_PAN_SPEED;
    }
    // UP
    if (KEY_MAP[38]) {
        camera.center.y -= 1 * CAMERA_PAN_SPEED;
    }
    // DOWN
    if (KEY_MAP[40]) {
        camera.center.y += 1 * CAMERA_PAN_SPEED;
    }

    // W - zoom in
    if (KEY_MAP[87]) {
        camera.zoomLevel = Math.min(4, camera.zoomLevel + 0.1);
    }
    // S - zoom out
    if (KEY_MAP[83]) {
        camera.zoomLevel = Math.max(1, camera.zoomLevel - 0.1);
    }
}
