import axios from 'axios'

// TODO: get base url as param
const axiosAPI = axios.create({
  baseURL: 'http://localhost:8080/api/v1/'
})

const apiRequest = (method, url, data) => {
  const headers = {
    authorization: ''
  }

  return axiosAPI({
    method,
    url,
    data: data,
    headers
  }).then(res => {
    return Promise.resolve(res.data)
  })
    .catch(err => {
      return Promise.reject(err)
    })
}

export const getReq = (url, data) => apiRequest('get', url, data)
export const postReq = (url, data) => apiRequest('post', url, data)
export const deleteReq = (url, data) => apiRequest('delete', url, data)
