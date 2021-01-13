import { getReq, postReq, SliceData } from '../../common/utils/api'

export const getShopsSlice = async (data: SliceData) => {
  return await getReq('shops/slice', data)
}
