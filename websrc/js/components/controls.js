export const KEY_MAP = {};

export function initControls() {
    console.log('controls initialized');
    document.addEventListener('keydown', key => {
        KEY_MAP[key.keyCode] = true;
    });

    document.addEventListener('keyup', key => {
        KEY_MAP[key.keyCode] = false;
    });
}
