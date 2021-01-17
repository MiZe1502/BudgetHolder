import {getReq, SliceData, postReq, EntityData} from '../../common/utils/api'
import { Category } from './types'

export const getCategoriesTree = async () => {
  return await getReq('categories/tree', {})
}

export const getCategoriesList = async () => {
  return await getReq('categories/list', {})
}

export const addNewCategory = async (data: Category) => {
  return await postReq('categories/new', data)
}

export const removeCategory = async (data: EntityData) => {
  return await postReq('categories/remove', data)
}

export const updateCategory = async (data: Category) => {
  return await postReq('categories/update', data)
}
