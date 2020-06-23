import { writable } from 'svelte/store'

export const yandexMapsReady = writable(false);
export const googleMapsReady = writable(false);