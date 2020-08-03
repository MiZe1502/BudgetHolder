import {get, writable} from 'svelte/store'
import {Purchase} from "../pages/Budget/types";
import {LoadingStatus} from "./utils";
import {generateNewArtificialId} from "./common";

export const purchases = writable<Purchase[]>([]);
export const purchasesTotal = writable<number>(0);
export const purchasesStatus = writable<LoadingStatus>(LoadingStatus.None);

export const removePurchaseFromStore = (id: number): void => {
    purchases.update((purchases) => purchases.filter((purchase) => purchase.id !== id))
    purchasesTotal.update(total => total - 1);
}

export const removeGoodsItemDetailsFromPurchase = (detailsId: number, purchaseId: number): void => {
    purchases.update((purchases) => {
        const currentPurchase = purchases.find((purchase) => purchase.id === purchaseId);
        currentPurchase.goods = currentPurchase.goods.filter((goodsItem) => goodsItem.id !== detailsId);

        console.log(currentPurchase)
        return purchases;
    })
}

export const addPurchaseToStore = (newPurchase: Purchase) => {
    newPurchase.id = generateNewArtificialId(purchases);

    purchases.update((purchases) => {
        return [...purchases, newPurchase];
    })
    purchasesTotal.update(total => total + 1);
}
