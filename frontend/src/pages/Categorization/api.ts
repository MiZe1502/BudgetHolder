import { getReq, SliceData, postReq } from '../../common/utils/api'
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
