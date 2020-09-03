export const addDataToLocalStorage = (key: string, data: any): void => {
    localStorage.setItem(key, JSON.stringify(data));
}

export const getDataFromLocalStorageByKey = (key: string): any => {
    return JSON.parse(localStorage.getItem(key));
}

export const removeDataFromLocalStorageByKey = (key: string): void => {
    localStorage.removeItem(key);
}