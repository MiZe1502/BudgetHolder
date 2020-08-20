import {writable} from "svelte/store";

interface UserData {
    name?: string;
    surname?: string;
    login: string;
}

interface UserSession {
    userData: UserData;
    lastLogin: Date;
}

export const authStatus = writable<boolean>(false);

export const setAuthStatus = (isAuthorized: boolean) => {
    authStatus.set(isAuthorized);
}