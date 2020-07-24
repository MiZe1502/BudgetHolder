import {get, writable} from 'svelte/store'
import {Category, SimpleCategory} from "../pages/Categorization/types";
import {LoadingStatus} from "./utils";

export const categories = writable<Category[]>([]);
export const categoriesTotal = writable<number>(0);
export const categoriesStatus = writable<LoadingStatus>(LoadingStatus.None);

export const simpleCategories = writable<SimpleCategory[]>([]);
export const simpleCategoriesStatus = writable<LoadingStatus>(LoadingStatus.None);

export const addCategoryToStore = (newCategory: Category) => {
    //TODO: remove when backend ready. artificial id generation
    let maxId = 0;
    for (let i = 0; i < get(categories).length; i++) {
        if (get(categories)[i].id > maxId) {
            maxId = get(categories)[i].id
        }
    }
    newCategory.id = maxId + 1;

    categories.update((categories) => {
        if (!newCategory.parentId) {
            return [...categories, newCategory]
        }
        const parentCategory = findCategoryById(newCategory.parentId, categories);

        if (parentCategory){
            parentCategory.categories = parentCategory.categories ? [...parentCategory.categories, newCategory] : [newCategory];
        } else {
            categories = [...categories, newCategory];
        }

        return categories;
    });
}

const findCategoryById = (id: number, categories: Category[]): Category => {
    let category = null;

    for (let i = 0; i < categories.length; i++) {
        if (categories[i].id === id) {
            return categories[i];
        }

        if (categories[i].categories && categories[i].categories.length > 0) {
            category = findCategoryById(id, categories[i].categories);
            if (category) {
                return category;
            }
        }
    }

    return category;
}

export const removeCategoryFromStore = (id: number): void => {
    categories.update((categories) => categories.filter((category: Category) => category.id !== id));
    categoriesTotal.update(total => total - 1);
}

export const updateCategoryInStore = (updatedCategory: Category) => {
    categories.update((categories) => {
        //let categoryFromStore: Category = categories.find((category: Category) => category.id === updatedCategory.id)
        let categoryFromStore: Category = findCategoryById(updatedCategory.id, categories)

        if (categoryFromStore.parentId === updatedCategory.parentId) {
            categoryFromStore = updatedCategory;
            return categories;
        }

        //const oldParentCategory = categories.find((category) => category.id === categoryFromStore.parentId);
        const oldParentCategory = findCategoryById(categoryFromStore.parentId, categories);
        //const newParentCategory = categories.find((category) => category.id === updatedCategory.parentId);
        const newParentCategory = findCategoryById(updatedCategory.parentId, categories);

        newParentCategory.categories = newParentCategory.categories ? [...newParentCategory.categories, updatedCategory] : [updatedCategory];
        if (oldParentCategory) {
            oldParentCategory.categories = oldParentCategory.categories.filter((category) => category.id !== updatedCategory.id);
        } else {
            categories = categories.filter((category) => category.id !== updatedCategory.id)
        }

        return categories
    });
}