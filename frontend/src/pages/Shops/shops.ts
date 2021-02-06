import { derived, get, writable } from 'svelte/store'
import { Shop } from './types'
import { LoadingStatus } from '../../stores/utils'
import {
  addShop,
  getShop,
  getShopsSlice,
  removeShop,
  updateShop,
  getSimpleShopsList
} from './api'
import {
  convertSimpleData,
  ErrorResponse, SimpleReqDataItem,
  SuccessResponse,
  TableDataResponse
} from '../../common/utils/api'
import { SimpleDataItem } from '../Categorization/types'

export const shops = writable<Shop[]>([])
export const simpleShops = writable<SimpleDataItem[]>([])
export const allShops = writable<Shop[]>([])
export const shopsTotal = writable<number>(0)
export const shopsStatus = writable<LoadingStatus>(LoadingStatus.None)

export const getSimpleShopsData = async (name: string) => {
  await getSimpleShopsList()
    .then((res: SuccessResponse) => {
      const data = convertSimpleData(res.data as SimpleReqDataItem[])
      simpleShops.set(data)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const removeShopFromStore = async (id: number) => {
  await removeShop({ id })
    .then((res: SuccessResponse) => {
      shops.update((shops) => shops.filter((shop: Shop) => shop.id !== Number(res.message)))
      shopsTotal.update((total) => total - 1)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const addShopToStore = async (newShop: Shop) => {
  shopsStatus.set(LoadingStatus.Loading)

  await addShop(newShop)
    .then((res: SuccessResponse) => {
      const newId = Number(res.message)
      newShop.id = newId
      shops.update((shops) => {
        return [newShop, ...shops.slice(0, shops.length)]
      })
      shopsTotal.update((total) => total + 1)
      shopsStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      shopsStatus.set(LoadingStatus.Error)
    })
}

export const updateShopInStore = async (updatedShop: Shop) => {
  await updateShop(updatedShop)
    .then((res: SuccessResponse) => {
      shops.update((shops) => {
        const shopIndex = shops.findIndex((shop: Shop) => shop.id === updatedShop.id)
        shops[shopIndex] = res.data as Shop
        return shops
      })
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const updateCurrentShopsSlice = async (from: number, count: number) => {
  shopsStatus.set(LoadingStatus.Loading)

  await getShopsSlice({ from, count })
    .then((res: SuccessResponse) => {
      const data = res.data as TableDataResponse
      shops.set(data.data as Shop[])
      shopsTotal.set(data.total)
      shopsStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      shopsStatus.set(LoadingStatus.Error)
    })
}

export const getShopById = async (id: number) => {
  return await getShop({ id })
    .then((res: SuccessResponse) => {
      return res.data as Shop
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}
