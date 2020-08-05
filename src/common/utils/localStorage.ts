export const addDataToLocalStorage = (key: string, data: any): void => {
    localStorage.setItem(key, JSON.stringify(data));
}
