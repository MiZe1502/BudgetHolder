import {writable} from "svelte/store";
import {mockedUsers} from "../pages/Auth/data";

export interface UserData {
    name?: string;
    surname?: string;
    login: string;
    image?: string;
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

export const authStatus = writable<boolean>(false);

export const currentAuthData = writable<AuthData>({} as AuthData);
export const currentRegData = writable<RegistrationData>({} as RegistrationData);

export const currentSession = writable<UserSession>({} as UserSession);

export const setAuthStatus = (isAuthorized: boolean) => {
    authStatus.set(isAuthorized);
}

export const setCurrentSession = (userData: UserData) => {
    currentSession.set({
        lastLogin: new Date(),
        userData: userData,
    })
}

//TODO: mock method to emulate auth
export const mockAuthorize = (authData: AuthData): string | undefined => {
    const currentUser = mockedUsers.find((item) => item.login === authData.login);

    if (!currentUser) {
        return `User with login ${authData.login} not found`;
    }

    if (currentUser.password !== authData.password) {
        return `Incorrect password for user ${authData.login}`;
    }

    setCurrentSession(currentUser);
    setAuthStatus(true);
    return;
}