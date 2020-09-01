import {get, writable, derived, Writable} from 'svelte/store'
import {GoodsData, GoodsDetails, GoodsSuggestion, GoodsItem} from "../pages/Budget/types";
import {LoadingStatus} from "./utils";
import {generateNewArtificialId} from "./common";

export const goods = writable<GoodsData[]>([]);
export const allGoods = writable<GoodsData[]>([]);
export const goodsTotal = writable<number>(0);
export const goodsStatus = writable<LoadingStatus>(LoadingStatus.None);

export const goodsForSuggestions = derived<Writable<GoodsData[]>, GoodsSuggestion[]>(allGoods, ($allGoods) => {
    return $allGoods.map((item) => ({
        id: item.id,
        value: item.name,
    }))
})

export const getGoodsItemById = (id: number): GoodsData => {
    return get(allGoods).find((item: GoodsData) => item.id === id);
}

export const addGoodsToStore = (newGoods: GoodsDetails[]) => {
    newGoods.forEach((item) => {
        if (item.id) {
            return;
        }

        const newId = generateNewArtificialId(allGoods)
        item.id = newId;
        const goodsItem = GoodsItem.fromGoodsDetails(item);

        allGoods.update((goods) => [...goods, goodsItem as GoodsData]);
        goodsTotal.update((total) => total + 1);
    });
}

export const removeGoodsFromStore = (id: number) => {
    goods.update((goods) => goods.filter((goodsItem: GoodsData) => goodsItem.id !== id));
    allGoods.update((goods) => goods.filter((goodsItem: GoodsData) => goodsItem.id !== id));
    goodsTotal.update((total) => total - 1);
}

export const updateGoodsItemInStore = (updatedGoodsItem: GoodsData) => {
    goods.update((goods) => {
        let goodsitemFromStore: GoodsData = goods.find((goodsItem) => goodsItem.id === updatedGoodsItem.id);
        goodsitemFromStore = updatedGoodsItem;
        return goods;
    });
    allGoods.update((goods) => {
        let goodsitemFromStore: GoodsData = goods.find((goodsItem) => goodsItem.id === updatedGoodsItem.id);
        goodsitemFromStore = updatedGoodsItem;
        return goods;
    });
}

export const updateCurrentGoodsSlice = (from: number, to: number) => {
    goods.set([...get(allGoods).slice(from, to + 1)]);
}
