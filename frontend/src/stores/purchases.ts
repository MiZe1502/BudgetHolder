import { get, writable } from 'svelte/store'
import { GoodsDetails, Purchase, ValidationResult } from '../pages/Budget/types'
import { LoadingStatus } from './utils'
import { generateNewArtificialId } from './common'
import { v4 as uuidv4 } from 'uuid'
import { Shop } from '../pages/Shops/types'
import { SimpleDataItem } from '../pages/Categorization/types'
import { removeDataFromLocalStorageByKey } from '../common/utils/localStorage'

export const purchaseLocalStorageKey = 'CURRENT_PURCHASE_FORM_STATE'
export const purchaseLocalStorageUpdateInterval = 20000

export const currentPurchase = writable<Purchase>(new Purchase())

export const validationResults = writable<ValidationResult[]>([])

export const purchases = writable<Purchase[]>([])
export const purchasesTotal = writable<number>(0)
export const purchasesStatus = writable<LoadingStatus>(LoadingStatus.None)

export const removePurchaseFromStore = (id: number): void => {
  purchases.update((purchases) => purchases.filter((purchase) => purchase.id !== id))
  purchasesTotal.update(total => total - 1)
}

export const removeGoodsItemDetailsFromPurchase = (detailsId: number, purchaseId: number): void => {
  purchases.update((purchases) => {
    const currentPurchase = purchases.find((purchase) => purchase.id === purchaseId)
    currentPurchase.goods = currentPurchase.goods.filter((goodsItem) => goodsItem.id !== detailsId)

    console.log(currentPurchase)
    return purchases
  })
}

export const addPurchaseToStore = (newPurchase: Purchase) => {
  newPurchase.id = generateNewArtificialId(purchases)

  purchases.update((purchases) => {
    return [...purchases, newPurchase]
  })
  purchasesTotal.update(total => total + 1)

  clearCurrentPurchaseData()
}

export const clearCurrentPurchaseData = () => {
  currentPurchase.set(new Purchase())
  removeDataFromLocalStorageByKey(purchaseLocalStorageKey)
}

export const updateValidationResults = (message: string, counter?: number) => {
  const validationResult: ValidationResult = {
    goodsItemCounter: counter + 1,
    message: message
  }

  validationResults.update((items) => [...items, validationResult])
}

export const clearValidationResults = () => {
  validationResults.set([])
}

export const updatePurchaseDataInStore = (newData: Purchase) => {
  purchases.update((purchases) => {
    const purchase = purchases.find((item) => item.id === newData.id)
    purchase.comment = newData.comment
    purchase.totalPrice = newData.totalPrice
    purchase.shop = newData.shop
    purchase.date = newData.date

    return purchases
  })
}

export const updatePurchaseGoodsItemInStore = (newData: GoodsDetails) => {
  purchases.update((purchases) => {
    const purchase = purchases.find((item) => item.id === newData.purchaseId)
    const goodsItem = purchase.goods.find((item) => item.id === newData.id)
    goodsItem.price = newData.price
    goodsItem.comment = newData.comment
    goodsItem.category = newData.category
    goodsItem.amount = newData.amount

    return purchases
  })
}
