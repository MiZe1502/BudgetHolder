import { get, writable } from 'svelte/store'
import {
  addDataToLocalStorage,
  removeDataFromLocalStorageByKey
} from '../../common/utils/localStorage'
import {
  authReq,
  editUserReq,
  getUserReq,
  logoutReq,
  registrationReq
} from './api'
import { ErrorResponse, SuccessResponse } from '../../common/utils/api'
import { parseJwt } from '../../common/utils/utils'

export interface UserData {
    name?: string;
    surname?: string;
    login: string;
    image?: string;
}

export interface UserDataExtended extends UserData{
    password: string;
    newPassword: string;
}

export interface UserSession {
    userData: UserData;
    lastLogin: Date;
}

export interface AuthData {
    login: string;
    password: string;
}

export interface RegistrationData extends AuthData {
    login: string;
    password: string;
    name?: string;
    surname?: string;
    image?: string;
}

export const sessionKey = 'SESSION_KEY'
export const authStatusKey = 'AUTH_STATUS_KEY'
export const currentUser = 'CURRENT_USER'

export const authStatus = writable<boolean>(false)

export const currentAuthData = writable<AuthData>({} as AuthData)
export const currentRegData = writable<RegistrationData>({} as RegistrationData)

export const currentSession = writable<UserSession>({} as UserSession)
const currentTokenExpirationTimeout = writable<number | null>(null)

export const setAuthStatus = (isAuthorized: boolean) => {
  authStatus.set(isAuthorized)
}

export const setCurrentSession = (userData: UserData) => {
  currentSession.set({
    lastLogin: new Date(),
    userData: userData
  })
}

export const authorize = async (authData: AuthData) => {
  const res = await authReq(authData)

  if (!res.status) {
    return 'Authorization error'
  }

  updateToken((res as SuccessResponse).message)
}

export const updateToken = (token: string) => {
  if (!token) {
    return
  }

  addDataToLocalStorage(sessionKey, token)

  if (get(currentTokenExpirationTimeout)) {
    clearTimeout(get(currentTokenExpirationTimeout))
    return
  }

  const parsedToken: {SessionID: string, exp: number} | null = parseJwt(token)
  if (!parsedToken) {
    return
  }
  const expirationTimeout = countDiffToTokenExpiration(parsedToken.exp)
  currentTokenExpirationTimeout.set(setTimeout(async () => {
    await logout()
  }, expirationTimeout))
}

const countDiffToTokenExpiration = (expirationDateTime: number) => {
  return (new Date(expirationDateTime * 1000)).getTime() - Date.now()
}

export const logout = async () => {
  await logoutReq()
    .catch((err: ErrorResponse) => {
      console.log(err)
    })

  setAuthStatus(false)
  setCurrentSession({} as UserData)

  removeDataFromLocalStorageByKey(sessionKey)
  removeDataFromLocalStorageByKey(authStatusKey)
}

export const getUserData = async (login: string) => {
  const res = await getUserReq({ login })

  if (!res.status) {
    return 'Error fetching user data'
  }

  setCurrentSession((res as SuccessResponse).data as UserData)
  addDataToLocalStorage(currentUser, (res as SuccessResponse).data)

  setAuthStatus(true)
  addDataToLocalStorage(authStatusKey, get(authStatus))
}

export const registerUser = async (regData: RegistrationData) => {
  const res = await registrationReq(Object.assign({ group_name: 'test_group' }, regData))

  if (!res.status) {
    return 'Error while new user registration'
  }

  const err = await authorize({ login: regData.login, password: regData.password })
  if (err) {
    return err
  }

  setCurrentSession(Object.assign({}, regData, { password: undefined }))
  addDataToLocalStorage(currentUser, Object.assign({}, regData, { password: undefined }))

  setAuthStatus(true)
  addDataToLocalStorage(authStatusKey, get(authStatus))
}

export const clearAuthAndRegData = () => {
  currentAuthData.set({} as AuthData)
  currentRegData.set({} as RegistrationData)
}

export const updateUserDataInStore = async (newData: UserDataExtended, login: string) => {
  await editUserReq(newData)
    .then((res: SuccessResponse) => {
      addDataToLocalStorage(currentUser, newData)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}
