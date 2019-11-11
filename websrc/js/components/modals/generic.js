const MODAL_MAPPING = {
    tileStats: 'js-tile-info-cntr'
};
export function openModal(modalName) {
    document
        .querySelector('.' + MODAL_MAPPING[modalName])
        .classList.remove('closed');
}

export function addGenericModalListeners() {
    let tEl = document.querySelector('.js-tile-info-cntr');
    tEl.addEventListener('click', e => {
        e.stopPropagation();
        tEl.classList.add('closed');
    });
}
