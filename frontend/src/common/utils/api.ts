import axios from 'axios'
import { getDataFromLocalStorageByKey } from './localStorage'
import { sessionKey } from '../../stores/auth'

// TODO: get base url as param
const axiosAPI = axios.create({
  baseURL: 'http://localhost:8080/api/v1/'
})

const apiRequest = (method, url, data = {}, params = {}): Promise<SuccessResponse | ErrorResponse> => {
  const headers = {
    authorization: ''
  }

  const token = getDataFromLocalStorageByKey(sessionKey)

  if (token) {
    headers.authorization = token
  }

  return axiosAPI({
    method,
    url,
    data,
    params,
    headers
  }).then(res => {
    const data: SuccessResponse = res.data

    if (data.data) {
      data.data = JSON.parse(data.data as string)
    }

    return Promise.resolve(data)
  })
    .catch(err => {
      console.log(err)
      return Promise.resolve(err)
    })
}

export const getReq = (url, params) => apiRequest('get', url, {}, params)
export const postReq = (url, data) => apiRequest('post', url, data)
export const deleteReq = (url, data) => apiRequest('delete', url, data)

export interface SimpleResponse {
  status: boolean;
}

export interface TableDataResponse {
  total: number;
  data: Record<string, any>
}

export interface SuccessResponse extends SimpleResponse {
  message: string
  data: string | Record<string, any>
}

export interface ErrorResponse extends SimpleResponse {
  errorCode: number
}

export interface SliceData {
  from: number;
  count: number;
}
