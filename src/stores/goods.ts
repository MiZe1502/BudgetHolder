import {get, writable} from 'svelte/store'
import {GoodsData} from "../pages/Budget/types";
import {LoadingStatus} from "./utils";

export const goods = writable<GoodsData[]>([]);
export const goodsTotal = writable<number>(0);
export const goodsStatus = writable<LoadingStatus>(LoadingStatus.None);

