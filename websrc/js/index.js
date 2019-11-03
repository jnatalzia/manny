import '../css/main.scss';
import { ctx } from './components/ctx';
import { update } from './gameloop';
import { initControls } from './components/controls';

window.addEventListener('load', () => {
    initControls();
    update(ctx);
});
