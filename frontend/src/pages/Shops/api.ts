import {EntityData, getReq, postReq, SliceData} from '../../common/utils/api'

export const getShopsSlice = async (data: SliceData) => {
  return await getReq('shops/slice', data)
}

export const getShop = async (data: EntityData) => {
  return await getReq('shops/get', data)
}

export const removeShop = async (data: EntityData) => {
  return await postReq('shops/remove', data)
}