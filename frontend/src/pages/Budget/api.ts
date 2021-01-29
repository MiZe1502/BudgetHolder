import { EntityData, getReq, postReq, SliceData } from '../../common/utils/api'
import { GoodsData } from './types'

export const getGoodsItemsSlice = async (data: SliceData) => {
  return await getReq('goods/slice', data)
}

export const removeGoodsItem = async (data: EntityData) => {
  return await postReq('goods/remove', data)
}

export const updateGoodsItem = async (data: GoodsData) => {
  return await postReq('goods/update', data)
}

export const addGoodsItem = async (data: GoodsData) => {
  return await postReq('goods/new', data)
}

export const getPurchaseWithGoodsSlice = async (data: SliceData) => {
  return await getReq('purchases/slice', data)
}

export const removePurchase = async (data: EntityData) => {
  return await postReq('purchases/remove', data)
}
