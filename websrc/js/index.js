import '../css/main.scss';
import { ctx } from './components/ctx';
import { update } from './gameloop';
import { initControls } from './components/controls';
import { initGamestate } from './components/gamestate';

window.addEventListener('load', () => {
    initControls();

    initGamestate().then(() => {
        update(ctx);
    });
});
