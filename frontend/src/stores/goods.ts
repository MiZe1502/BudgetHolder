import { derived, get, writable, Writable } from 'svelte/store'
import {
  GoodsData,
  GoodsDetails,
  GoodsItem,
  GoodsSuggestion
} from '../pages/Budget/types'
import { LoadingStatus } from './utils'
import { generateNewArtificialId } from './common'
import {
  addGoodsItem,
  getGoodsItemsSlice,
  removeGoodsItem, updateGoodsItem
} from '../pages/Budget/api'
import {
  ErrorResponse,
  SuccessResponse,
  TableDataResponse
} from '../common/utils/api'

export const goods = writable<GoodsData[]>([])
export const allGoods = writable<GoodsData[]>([])
export const goodsTotal = writable<number>(0)
export const goodsStatus = writable<LoadingStatus>(LoadingStatus.None)

export const goodsForSuggestions = derived<Writable<GoodsData[]>, GoodsSuggestion[]>(allGoods, ($allGoods) => {
  return $allGoods.map((item) => ({
    id: item.id,
    value: item.name
  }))
})

export const getGoodsItemById = (id: number): GoodsData => {
  return get(allGoods).find((item: GoodsData) => item.id === id)
}

export const addGoodsItemToStore = async (newGoodsItem: GoodsData) => {
  goodsStatus.set(LoadingStatus.Loading)
  await addGoodsItem(newGoodsItem)
    .then((res: SuccessResponse) => {
      const newId = Number(res.message)
      newGoodsItem.id = newId
      goods.update((goods) => {
        return [newGoodsItem, ...goods.slice(0, goods.length)]
      })
      goodsTotal.update((total) => total + 1)
      goodsStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      goodsStatus.set(LoadingStatus.Error)
    })
}

export const addGoodsToStore = async (newGoods: GoodsDetails[]) => {
  newGoods.forEach((item) => {
    if (item.id) {
      return
    }

    const newId = generateNewArtificialId(allGoods)
    item.id = newId
    const goodsItem = GoodsItem.fromGoodsDetails(item)

    allGoods.update((goods) => [...goods, goodsItem as GoodsData])
    goodsTotal.update((total) => total + 1)
  })
}

export const removeGoodsFromStore = async (id: number) => {
  await removeGoodsItem({ id })
    .then((res: SuccessResponse) => {
      goods.update((goods) => goods.filter((goodsItem: GoodsData) => goodsItem.id !== Number(res.message)))
      goodsTotal.update((total) => total - 1)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const updateGoodsItemInStore = async (updatedGoodsItem: GoodsData) => {
  await updateGoodsItem(updatedGoodsItem)
    .then((res: SuccessResponse) => {
      goods.update((goods) => {
        let goodsitemFromStore: GoodsData = goods.find((goodsItem) => goodsItem.id === updatedGoodsItem.id)
        goodsitemFromStore = updatedGoodsItem
        return goods
      })
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const updateCurrentGoodsSlice = async (from: number, count: number) => {
  goodsStatus.set(LoadingStatus.Loading)

  await getGoodsItemsSlice({ from, count })
    .then((res: SuccessResponse) => {
      const data = res.data as TableDataResponse
      goods.set(data.data as GoodsData[])
      goodsTotal.set(data.total)
      goodsStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      goodsStatus.set(LoadingStatus.Error)
    })
}
