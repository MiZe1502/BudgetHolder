import {getReq, postReq} from '../../common/utils/api'
import {AuthData, RegistrationData, UserData} from '../../stores/auth'

export const authReq = async (data: AuthData) => {
  return await postReq('user/auth', data)
}

export const logoutReq = async (login: string, password: string) => {
  return await postReq('user/logout', { login, password })
}

export const registrationReq = async (data: RegistrationData) => {
  return await postReq('user/new', data)
}

export const getUserReq = async (data: UserData) => {
  return await getReq('user/full', data)
}