
export const getCurrentPath = ():string => {
  return window.location.pathname
}

export const parseJwt = (token: string) => {
  try {
    return JSON.parse(atob(token.split('.')[1]))
  } catch (e) {
    return null
  }
}
