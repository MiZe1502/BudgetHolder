import { derived, get, writable, Writable } from 'svelte/store'
import {
  GoodsData,
  GoodsDetails,
  GoodsItem,
  GoodsSuggestion
} from '../pages/Budget/types'
import { LoadingStatus } from './utils'
import { generateNewArtificialId } from './common'
import { getGoodsItemsSlice } from '../pages/Budget/api'
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

export const addGoodsToStore = (newGoods: GoodsDetails[]) => {
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

export const removeGoodsFromStore = (id: number) => {
  goods.update((goods) => goods.filter((goodsItem: GoodsData) => goodsItem.id !== id))
  allGoods.update((goods) => goods.filter((goodsItem: GoodsData) => goodsItem.id !== id))
  goodsTotal.update((total) => total - 1)
}

export const updateGoodsItemInStore = (updatedGoodsItem: GoodsData) => {
  goods.update((goods) => {
    let goodsitemFromStore: GoodsData = goods.find((goodsItem) => goodsItem.id === updatedGoodsItem.id)
    goodsitemFromStore = updatedGoodsItem
    return goods
  })
  allGoods.update((goods) => {
    let goodsitemFromStore: GoodsData = goods.find((goodsItem) => goodsItem.id === updatedGoodsItem.id)
    goodsitemFromStore = updatedGoodsItem
    return goods
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
    // goods.set([...get(allGoods).slice(from, to + 1)])
}
