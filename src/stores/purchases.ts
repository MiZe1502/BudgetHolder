import {get, writable} from 'svelte/store'
import {Purchase} from "../pages/Budget/types";
import {LoadingStatus} from "./utils";

export const purchases = writable<Purchase[]>([]);
export const purchasesTotal = writable<number>(0);
export const purchasesStatus = writable<LoadingStatus>(LoadingStatus.None);
