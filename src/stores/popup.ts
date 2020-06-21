import { writable } from 'svelte/store';
import { get } from 'svelte/store';
import { v4 as uuidv4 } from 'uuid';

//TODO: We can implement this logic throught linked list

export const defaultPopupZIndex = 10;

export interface PopupState {
    zIndex?: number;
    uuid?: string;
}
type PopupRecord = Record<string, PopupState>;

export const openedPopups = writable({});

export const findCurrentTopPopup = (): PopupState => {
    const popups = get(openedPopups) as PopupRecord;

    let topPopup: PopupState = {zIndex: defaultPopupZIndex};
    let topZIndex = defaultPopupZIndex;
    for (let popupGuid in popups) {
        if (popups[popupGuid].zIndex > topZIndex) {
            topZIndex = popups[popupGuid].zIndex;
            topPopup = popups[popupGuid];
        } 
    }

    return topPopup;
}

export const addPopupToState = (): PopupState => {
    const uuid = uuidv4();

    const topPopup = findCurrentTopPopup();
    const curPopup = {zIndex: topPopup.zIndex + 1, uuid: uuid}
    openedPopups.set({...get(openedPopups), [uuid]: curPopup})

    return curPopup;
}

export const movePopupToTheTop = (uuid: string): PopupState => {
    let newTopPopup = {};
    openedPopups.update((store: PopupRecord) => {
        const topPopup = findCurrentTopPopup();
        newTopPopup = {
            zIndex: topPopup.zIndex + 1, 
            uuid: uuid
        }
        return {
            ...store,
            [uuid]: newTopPopup,
        }
    })
    return newTopPopup;
}

// export const removePopupFromStore = (uuid: string) => {
//     openedPopups.update((store: PopupRecord) => {
//         const {[uuid], ...rest} = store;
//         return {
//             ...store,
//             [uuid]: newTopPopup,
//         }
//     })
// }