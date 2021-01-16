import { getReq, SliceData } from '../../common/utils/api'

export const getCategoriesTree = async () => {
  return await getReq('categories/tree', {})
}

export const getCategoriesList = async () => {
  return await getReq('categories/list', {})
}
