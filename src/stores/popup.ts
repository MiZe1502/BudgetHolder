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

export const openedPopups = writable([]);

export const findCurrentTopPopup = (): PopupState => {
    const popups = get(openedPopups) as PopupRecord[];

    let topPopup: PopupState = popups[0];

    for (let i = 0; i < popups.length; i++) {
        if (popups[i].zIndex > topPopup.zIndex) {
            topPopup = popups[i];
        }
    }

    return topPopup;
}

export const addPopupToState = (uuid: string): PopupState => {
    const topPopup = findCurrentTopPopup();

    let curPopup: PopupState = {
        zIndex: defaultPopupZIndex, 
        uuid: uuid
    };

    if (topPopup) {
        curPopup.zIndex = topPopup.zIndex + 1;
    }

    openedPopups.update((popups) => [...popups, curPopup])

    return curPopup;
}

export const movePopupToTheTop = (uuid: string): PopupState => {
    let newTopPopup: PopupState = {};

    openedPopups.update((popups) => {
        const topPopup = findCurrentTopPopup();

        newTopPopup = popups.find((popup) => popup.uuid === uuid);
        newTopPopup.zIndex = topPopup.zIndex + 1;
        
        return popups;
    })

    return newTopPopup;
}

export const removePopupFromStore = (uuid: string) => {
    openedPopups.update((popups) => popups.filter((popup) => popup.uuid !== uuid));
}