export let ctx;
export let canvas;

window.addEventListener('load', () => {
    canvas = document.querySelector('#canvas');
    ctx = canvas.getContext('2d');
});
