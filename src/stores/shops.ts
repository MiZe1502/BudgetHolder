import { writable, get } from 'svelte/store'
import { Shop } from '../pages/Shops/types'
import { LoadingStatus } from './utils'

export const shops = writable<Shop[]>([])
export const allShops = writable<Shop[]>([]);
export const shopsTotal = writable<number>(0)
export const shopsStatus = writable<LoadingStatus>(LoadingStatus.None);

export const removeShopFromStore = (id: number) => {
    shops.update((shops) => shops.filter((shop: Shop) => shop.id !== id));
    allShops.update((shops) => shops.filter((shop: Shop) => shop.id !== id));
    shopsTotal.update((total) => total - 1);
}

export const addShopToStore = (newShop: Shop) => {
    let maxId = 0;
    for (let i = 0; i < get(allShops).length; i++) {
        if (get(allShops)[i].id > maxId) {
            maxId = get(allShops)[i].id
        }
    }
    newShop.id = maxId + 1;


    allShops.update((shops) => [...shops, newShop]);
    shopsTotal.update((total) => total + 1);
}

export const updateShopInStore = (updatedShop: Shop) => {
    shops.update((shops) => {
        let shopFromStore: Shop = shops.find((shop: Shop) => shop.id === updatedShop.id)
        shopFromStore = updatedShop;
        return shops;
    });
    allShops.update((shops) => {
        let shopFromStore: Shop = shops.find((shop: Shop) => shop.id === updatedShop.id)
        shopFromStore = updatedShop;
        return shops;
    });
}

export const updateCurrentShopsSlice = (from: number, to: number) => {
    shops.set([...get(allShops).slice(from, to + 1)]);
}