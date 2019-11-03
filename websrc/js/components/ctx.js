export let ctx;
export let canvas;

const FOOTER_HEIGHT = 75;
export const CANVAS_WIDTH = window.innerWidth;
export const CANVAS_HEIGHT = window.innerHeight - FOOTER_HEIGHT;
export const DEVICE_DPR = window.devicePixelRatio || 1;

window.addEventListener('load', () => {
    const dpr = DEVICE_DPR;
    canvas = document.querySelector('#canvas');
    canvas.style.width = `${CANVAS_WIDTH}px`;
    canvas.style.height = `${CANVAS_HEIGHT}px`;
    canvas.width = CANVAS_WIDTH * dpr;
    canvas.height = CANVAS_HEIGHT * dpr;

    ctx = canvas.getContext('2d');
    ctx.scale(dpr, dpr);
});
