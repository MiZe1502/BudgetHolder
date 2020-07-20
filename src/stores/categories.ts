import {writable} from 'svelte/store'
import {Category} from "../pages/Categorization/types";
import {LoadingStatus} from "./utils";
import {Shop} from "../pages/Shops/types";

export const categories = writable<Category[]>([]);
export const categoriesTotal = writable<number>(0);
export const categoriesStatus = writable<LoadingStatus>(LoadingStatus.None);

export const addCategoryToStore = (newCategory: Category) => {
    categories.update((categories) => [...categories, newCategory]);
}

export const removeCategoryFromStore = (id: number): void => {
    categories.update((shops) => shops.filter((shop: Shop) => shop.id !== id));
    categoriesTotal.update(total => total - 1);
}