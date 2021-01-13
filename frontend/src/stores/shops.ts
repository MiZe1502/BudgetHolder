import { writable, get, derived } from 'svelte/store'
import { Shop } from '../pages/Shops/types'
import { LoadingStatus } from './utils'
import { generateNewArtificialId } from './common'
import { getShopsSlice } from '../pages/Shops/api'
import { ErrorResponse, SuccessResponse } from '../common/utils/api'

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

export const removeShopFromStore = (id: number) => {
  shops.update((shops) => shops.filter((shop: Shop) => shop.id !== id))
  allShops.update((shops) => shops.filter((shop: Shop) => shop.id !== id))
  shopsTotal.update((total) => total - 1)
}

export const addShopToStore = (newShop: Shop) => {
  const newId = generateNewArtificialId(allShops)
  newShop.id = newId

  allShops.update((shops) => [...shops, newShop])
  shopsTotal.update((total) => total + 1)
}

export const updateShopInStore = (updatedShop: Shop) => {
  shops.update((shops) => {
    let shopFromStore: Shop = shops.find((shop: Shop) => shop.id === updatedShop.id)
    shopFromStore = updatedShop
    return shops
  })
  allShops.update((shops) => {
    let shopFromStore: Shop = shops.find((shop: Shop) => shop.id === updatedShop.id)
    shopFromStore = updatedShop
    return shops
  })
}

export const updateCurrentShopsSlice = async (from: number, count: number) => {
  shopsStatus.set(LoadingStatus.Loading)

  await getShopsSlice({ from, count })
    .then((res: SuccessResponse) => {
      shopsStatus.set(LoadingStatus.Finished)
      console.log(res.data)
      shops.set(res.data as Shop[])
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

// export const updateCurrentShopsSlice = (from: number, to: number) => {
//   shops.set([...get(allShops).slice(from, to + 1)])
// }

export const getShopById = (id: number): Shop => {
  return get(allShops).find((shop: Shop) => shop.id === id)
}
