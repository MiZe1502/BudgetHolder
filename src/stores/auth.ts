import {writable} from "svelte/store";

interface UserData {
    name?: string;
    surname?: string;
    login: string;
    image?: string;
}

interface UserSession {
    userData: UserData;
    lastLogin: Date;
}

interface AuthData {
    login: string;
    password: string;
}

export const authStatus = writable<boolean>(false);

export const currentAuthData = writable<AuthData>({} as AuthData);

export const setAuthStatus = (isAuthorized: boolean) => {
    authStatus.set(isAuthorized);
}