import { getReq, SliceData } from '../../common/utils/api'

export const getGoodsItemsSlice = async (data: SliceData) => {
  return await getReq('goods/slice', data)
}
