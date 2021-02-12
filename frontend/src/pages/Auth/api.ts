import { getReq, postReq } from '../../common/utils/api'
import { AuthData, RegistrationData, UserData } from './auth'

export const authReq = async (data: AuthData) => {
  return await postReq('user/auth', data)
}

export const logoutReq = async () => {
  return await postReq('user/logout', {})
}

export const registrationReq = async (data: RegistrationData) => {
  return await postReq('user/new', data)
}

export const getUserReq = async (data: UserData) => {
  return await getReq('user/full', data)
}

export const editUserReq = async (data: UserData) => {
  return await postReq('user/update', data)
}
