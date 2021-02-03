import { writable } from 'svelte/store'
import { GoodsDetails, Purchase, ValidationResult } from './types'
import { LoadingStatus } from '../../stores/utils'
import moment from 'moment'
import { removeDataFromLocalStorageByKey } from '../../common/utils/localStorage'
import {
  ErrorResponse,
  SuccessResponse,
  TableDataResponse
} from '../../common/utils/api'
import {
  addPurchase,
  getPurchaseWithGoodsSlice,
  removeGoodsDetailsFromPurchase,
  removePurchase,
  updateGoodsItemInPurchase,
  updatePurchase
} from './api'

export const dateTimeFormat = 'YYYY-MM-DDTHH:mm:ssZ'

export const purchaseLocalStorageKey = 'CURRENT_PURCHASE_FORM_STATE'
export const purchaseLocalStorageUpdateInterval = 20000

export const currentPurchase = writable<Purchase>(new Purchase())

export const validationResults = writable<ValidationResult[]>([])

export const purchases = writable<Purchase[]>([])
export const purchasesTotal = writable<number>(0)
export const purchasesStatus = writable<LoadingStatus>(LoadingStatus.None)

export const removePurchaseFromStore = async (id: number) => {
  await removePurchase({ id })
    .then((res: SuccessResponse) => {
      purchases.update((purchases) => purchases.filter((purchase) => purchase.id !== Number(res.message)))
      purchasesTotal.update(total => total - 1)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const removeGoodsItemDetailsFromPurchase = async (detailsId: number, purchaseId: number) => {
  await removeGoodsDetailsFromPurchase({ id: detailsId })
    .then((res: SuccessResponse) => {
      purchases.update((purchases) => {
        const currentPurchase = purchases.find((purchase) => purchase.id === purchaseId)
        currentPurchase.goods_data = currentPurchase.goods_data.filter((goodsItem) => goodsItem.id !== Number(res.message))
        return purchases
      })
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const addPurchaseToStore = async (newPurchase: Purchase) => {
  purchasesStatus.set(LoadingStatus.Loading)
  await addPurchase({ ...newPurchase, date: moment(new Date(newPurchase.date)).format(dateTimeFormat) })
    .then((res: SuccessResponse) => {
      const newId = Number(res.message)
      newPurchase.id = newId

      purchases.update((purchases) => {
        return [newPurchase, ...purchases.slice(0, purchases.length)]
      })
      purchasesTotal.update(total => total + 1)
      purchasesStatus.set(LoadingStatus.Finished)
      clearCurrentPurchaseData()
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      purchasesStatus.set(LoadingStatus.Error)
    })
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

export const updatePurchaseDataInStore = async (updatedPurchase: Purchase) => {
  await updatePurchase(updatedPurchase)
    .then((res: SuccessResponse) => {
      purchases.update((purchases) => {
        const purchaseIndex = purchases.findIndex((item) => item.id === updatedPurchase.id)

        purchases[purchaseIndex] = { ...res.data as Record<string, any>, goods_data: purchases[purchaseIndex].goods_data } as Purchase
        return purchases
      })
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const updatePurchaseGoodsItemInStore = async (newData: GoodsDetails) => {
  await updateGoodsItemInPurchase(newData)
    .then((res: SuccessResponse) => {
      purchases.update((purchases) => {
        const purchase = purchases.find((item) => item.id === newData.purchase_id)
        const goodsItem = purchase.goods_data.find((item) => item.id === newData.id)
        goodsItem.price = newData.price
        goodsItem.comment = newData.comment
        goodsItem.category = newData.category
        goodsItem.amount = newData.amount

        return purchases
      })
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const updatePurchasesWithGoodsSlice = async (from: number, count: number) => {
  purchasesStatus.set(LoadingStatus.Loading)
  await getPurchaseWithGoodsSlice({ from, count })
    .then((res: SuccessResponse) => {
      const data = res.data as TableDataResponse
      purchases.set(data.data as Purchase[])
      purchasesTotal.set(data.total)
      purchasesStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      purchasesStatus.set(LoadingStatus.Error)
    })
}
