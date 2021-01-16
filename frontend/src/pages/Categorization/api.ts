import { getReq, SliceData } from '../../common/utils/api'

export const getCategoriesTree = async () => {
  return await getReq('categories/tree', {})
}
