import { EntityData, getReq, postReq, SliceData } from '../../common/utils/api'
import {GoodsData, GoodsDetails, Purchase} from './types'

export const getGoodsItemsSlice = async (data: SliceData) => {
  return await getReq('goods/slice', data)
}

export const getSimpleGoodsForSuggestions = async (data: { name: string }) => {
  return await getReq('goods/list', data)
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

export const addPurchase = async (data: Purchase) => {
  return await postReq('purchases/new', data)
}

export const removePurchase = async (data: EntityData) => {
  return await postReq('purchases/remove', data)
}

export const updatePurchase = async (data: Purchase) => {
  return await postReq('purchases/update', data)
}

export const removeGoodsDetailsFromPurchase = async (data: EntityData) => {
  return await postReq('purchases/details/remove', data)
}

export const updateGoodsItemInPurchase = async (data: GoodsDetails) => {
  return await postReq('purchases/details/update', data)
}
