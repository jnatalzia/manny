import { currentGamestate } from '../gamestate';
import { DH_UNABLE_TO_CHECK_GENERATOR } from 'constants';
import { openModal } from './generic';

function renderPlayerTrader(t, container) {
    const traderEl = document.createElement('div');
    traderEl.innerText = `Trader ${t.id} (Owner: ${getNameForOwnerID(
        t.owner_id
    )}): Destination ${t.destination_id} | Carrying ${
        t.infected_crop
    } Infected Crop and ${t.safe_crop} Safe Crop | Cost ${t.trader_cost}`;
    container.appendChild(traderEl);
}

function renderGenericTrader(t, container) {
    const traderEl = document.createElement('div');
    traderEl.innerText = `Trader ${t.id}: Destination ${t.destination_id} | Cost ${t.trader_cost}`;
    container.appendChild(traderEl);
}

function renderPlayerDetails(tile, elements) {
    if (elements.crop) {
        const goodCropEl = document.createElement('div');
        goodCropEl.innerText = 'Good Crop: ';
        const badCropEl = document.createElement('div');
        badCropEl.innerText = 'Bad Crop: ';
        const goodCropSpanEl = document.createElement('span');
        goodCropSpanEl.classList.add('emphasized');
        goodCropSpanEl.innerText = tile.crop;
        const badCropSpanEl = document.createElement('span');
        badCropSpanEl.classList.add('emphasized');
        badCropSpanEl.innerText = tile.infected_crop;

        goodCropEl.appendChild(goodCropSpanEl);
        badCropEl.appendChild(badCropSpanEl);

        elements.crop.appendChild(goodCropEl);
        elements.crop.appendChild(badCropEl);
    }
}

function renderGenericTileInfo(tile, elements) {
    if (elements.sale) {
        let saleInfo = document.createElement('div');
        let saleEmp = document.createElement('span');
        saleInfo.innerText = 'Base Cost: ';
        saleEmp.innerText = tile.base_cost;
        saleInfo.appendChild(saleEmp);
        elements.sale.appendChild(saleInfo);
    }

    if (elements.traders) {
        let trds = tile.traders;
        trds.forEach(t => {
            t.current_player_owns
                ? renderPlayerTrader(t, elements.traders)
                : renderGenericTrader(t, elements.traders);
        });
    }
}

function getNameForOwnerID(oid) {
    let ownerID;
    switch (oid) {
        case currentGamestate.playerID:
            ownerID = 'You';
            break;
        case -1:
            ownerID = 'AI Farmer';
            break;
        default:
            ownerID = oid;
            break;
    }
    return ownerID;
}

function populateTileStats(tile) {
    let tileEl = document.querySelector('.js-tile-info');
    const cropEl = tileEl.querySelector('.js-crop');
    const tileOwner = tileEl.querySelector('.js-tile-owner');

    const elObj = {
        crop: cropEl,
        sale: tileEl.querySelector('.js-sale-info'),
        tileOwner,
        traders: tileEl.querySelector('.js-traders')
    };

    Object.keys(elObj).forEach(k => {
        elObj[k].innerHTML = '';
    });

    let ownerID = getNameForOwnerID(tile.owner_id);

    tileOwner.innerText = `Owner: ${ownerID}`;

    if (tile.current_player_owns) {
        renderPlayerDetails(tile, elObj);
    }

    renderGenericTileInfo(tile, elObj);
}

export function addTileStatListeners() {
    document
        .querySelector('.js-tile-info-modal')
        .addEventListener('click', e => {
            e.stopPropagation();
        });
}

export function openTileStats(tile) {
    openModal('tileStats');
    populateTileStats(tile);
}
