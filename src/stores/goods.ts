import {get, writable, derived, Writable} from 'svelte/store'
import {GoodsData, GoodsSuggestion} from "../pages/Budget/types";
import {LoadingStatus} from "./utils";

export const goods = writable<GoodsData[]>([]);
export const goodsTotal = writable<number>(0);
export const goodsStatus = writable<LoadingStatus>(LoadingStatus.None);

export const goodsForSuggestions = derived<Writable<GoodsData[]>, GoodsSuggestion[]>(goods, ($goods) => {
    return $goods.map((item) => ({
        id: item.id,
        value: item.name,
    }))
})

export const getGoodsItemById = (id: number): GoodsData => {
    return get(goods).find((item: GoodsData) => item.id === id);
}
