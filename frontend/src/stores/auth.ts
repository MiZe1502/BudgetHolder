import { get, writable } from 'svelte/store'
import { MockedUserData, mockedUsers } from '../pages/Auth/data'
import { addDataToLocalStorage } from '../common/utils/localStorage'
import { ValidationResult } from '../pages/Budget/types'
import { authReq, getUserReq, registrationReq } from '../pages/Auth/api'
import { SuccessResponse } from '../common/utils/api'

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

// TODO: mock store to save users
export const users = writable<MockedUserData[]>(mockedUsers)

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

  addDataToLocalStorage(sessionKey, (res as SuccessResponse).message)
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

export const updateUserDataInStore = (newData: UserDataExtended, login: string): string => {
  let error = ''

  users.update((users) => {
    const user = users.find((user) => login === user.login)

    if (newData.password !== user.password) {
      error = `Incorrect password for user ${user.login}`
    }

    if (newData.password === newData.newPassword) {
      error = 'You are not allowed to use the same password'
    }

    if (error) {
      return users
    }

    user.login = newData.login
    user.password = newData.newPassword
    user.name = newData.name
    user.surname = newData.surname

    return users
  })

  return error
}
