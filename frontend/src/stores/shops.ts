import { derived, get, writable } from 'svelte/store'
import { Shop } from '../pages/Shops/types'
import { LoadingStatus } from './utils'
import { generateNewArtificialId } from './common'
import {
  addShop,
  getShop,
  getShopsSlice,
  removeShop,
  updateShop
} from '../pages/Shops/api'
import {
  ErrorResponse,
  SuccessResponse,
  TableDataResponse
} from '../common/utils/api'

export const shops = writable<Shop[]>([])
export const allShops = writable<Shop[]>([])
export const shopsTotal = writable<number>(0)
export const shopsStatus = writable<LoadingStatus>(LoadingStatus.None)

export const simpleShops = derived(allShops, ($allShops) => {
  return $allShops.map((shop) => {
    return {
      id: shop.id,
      value: shop.name
    }
  })
})

export const getSimpleShopsData = () => {
  return get(allShops).map((shop: Shop) => {
    return {
      id: shop.id,
      value: shop.name
    }
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
}

export const updateShopInStore = async (updatedShop: Shop) => {
  await updateShop(updatedShop)
    .then((res: SuccessResponse) => {
      shops.update((shops) => {
        let shopFromStore: Shop = shops.find((shop: Shop) => shop.id === updatedShop.id)
        shopFromStore = updatedShop
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
      shopsStatus.set(LoadingStatus.Finished)
      shops.set(data.data as Shop[])
      shopsTotal.set(data.total)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
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
