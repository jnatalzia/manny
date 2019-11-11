import { canvas } from './ctx';

export const KEY_MAP = {};

export const ZOOM_IN_BTN_KEY = 'ZOOM_IN_BTN';
export const ZOOM_OUT_BTN_KEY = 'ZOOM_OUT_BTN';

const mouseDownEvts = ['touchstart', 'mousedown'];
const mouseUpEvts = ['touchend', 'mouseup'];
const mouseMoveEvts = ['touchmove', 'mousemove'];
let lastClick;

export function initControls() {
    initKeyControls();
    initUIControls();
    initTouchControls();
    console.log('controls initialized');
}

let mouseDiff = { x: 0, y: 0 };
let didClick = false;

export function getMouseMovement() {
    let mvment = Object.assign({}, mouseDiff);
    mouseDiff = { x: 0, y: 0 };
    return mvment;
}

export function didClickThisCycle() {
    return didClick;
}

export function getClickPoint() {
    didClick = false;
    return getPosFromEvt(lastClick);
}

function handleMapClick(e) {
    didClick = true;
    lastClick = e;
}

function initTouchControls() {
    const c = document.querySelector('#canvas');

    let lastMousePos = { x: -1, y: -1 };
    let dragging = false;
    let dragStartTime;

    const mouseDown = e => {
        e.stopPropagation();
        e.preventDefault();
        dragStartTime = Date.now();
        dragging = true;
        let { x, y } = getPosFromEvt(e);
        lastMousePos = { x, y };
    };

    const mouseUp = e => {
        e.stopPropagation();
        e.preventDefault();
        dragging = false;
        if (Date.now() - dragStartTime < 150) {
            // treat as click
            handleMapClick(e);
        }
        mouseDiff = { x: 0, y: 0 };
        lastMousePos = { x: -1, y: -1 };
    };

    const mouseMove = e => {
        if (!dragging) {
            return;
        }
        e.stopPropagation();
        e.preventDefault();

        let { x, y } = getPosFromEvt(e);

        // get the diff
        mouseDiff.x += x - lastMousePos.x;
        mouseDiff.y += y - lastMousePos.y;
        // reset the thing
        lastMousePos = { x, y };
    };

    mouseUpEvts.forEach(e => {
        c.addEventListener(e, mouseUp);
    });

    mouseDownEvts.forEach(e => {
        c.addEventListener(e, mouseDown);
    });

    mouseMoveEvts.forEach(e => {
        c.addEventListener(e, mouseMove);
    });
}

function getPosFromEvt(e) {
    let xPos;
    let yPos;
    if (e.touches) {
        let t = e.touches[0];
        xPos = t.screenX;
        yPos = t.screenY;
    } else {
        xPos = e.clientX;
        yPos = e.clientY;
    }

    return {
        x: xPos,
        y: yPos
    };
}

function initKeyControls() {
    document.addEventListener('keydown', key => {
        KEY_MAP[key.keyCode] = true;
    });

    document.addEventListener('keyup', key => {
        KEY_MAP[key.keyCode] = false;
    });
}

function initUIControls() {
    let zInEl = document.querySelector('.js-zoom-in');
    let zOutEl = document.querySelector('.js-zoom-out');

    mouseDownEvts.forEach(e => {
        zInEl.addEventListener(e, () => {
            KEY_MAP[ZOOM_IN_BTN_KEY] = true;
        });
        zOutEl.addEventListener(e, () => {
            KEY_MAP[ZOOM_OUT_BTN_KEY] = true;
        });
    });

    mouseUpEvts.forEach(e => {
        zInEl.addEventListener(e, () => {
            KEY_MAP[ZOOM_IN_BTN_KEY] = false;
        });

        zOutEl.addEventListener(e, () => {
            KEY_MAP[ZOOM_OUT_BTN_KEY] = false;
        });
    });
}
