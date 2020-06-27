import { writable } from 'svelte/store'

export const currentPage = writable(window.location.pathname)
