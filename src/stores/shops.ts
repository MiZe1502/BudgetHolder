import { writable } from 'svelte/store'
import { Shop } from '../pages/types'
import { LoadingStatus } from './utils'

export const shops = writable([])
export const shopsTotal = writable(0)
export const shopsStatus = writable(LoadingStatus.None);

export const removeShopFromStore = (id: number) => {
    shops.update((shops) => shops.filter((shop: Shop) => shop.id !== id));
    shopsTotal.update((total) => total - 1);
}

export const addShopToStore = (newShop: Shop) => {
    shops.update((shops) => [...shops, newShop]);
    shopsTotal.update((total) => total + 1);
}