import {get, writable} from 'svelte/store'

export const currentPage = writable(window.location.pathname)

//TODO: remove when backend ready. artificial id generation
export const generateNewArtificialId = (store): number => {
    let maxId = 0;
    for (let i = 0; i < get(store).length; i++) {
        if (get(store)[i].id > maxId) {
            maxId = get(store)[i].id
        }
    }

    return maxId + 1;
}

