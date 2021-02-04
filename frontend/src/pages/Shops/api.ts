import { EntityData, getReq, postReq, SliceData } from '../../common/utils/api'
import { Shop } from './types'

export const getShopsSlice = async (data: SliceData) => {
  return await getReq('shops/slice', data)
}

export const getShop = async (data: EntityData) => {
  return await getReq('shops/get', data)
}

export const removeShop = async (data: EntityData) => {
  return await postReq('shops/remove', data)
}

export const updateShop = async (data: Shop) => {
  return await postReq('shops/update', data)
}

export const addShop = async (data: Shop) => {
  return await postReq('shops/new', data)
}

export const getTopShopsSuggestions = async (data: {name: string}) => {
  return await getReq('shops/top', data)
}
