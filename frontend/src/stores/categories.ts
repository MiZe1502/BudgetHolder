import { get, writable } from 'svelte/store'
import {
  Category,
  SimpleDataItem
} from '../pages/Categorization/types'
import { LoadingStatus } from './utils'
import { generateNewArtificialId } from './common'
import {
  getCategoriesTree,
  getCategoriesList,
  addNewCategory
} from '../pages/Categorization/api'
import {
  ErrorResponse,
  SuccessResponse,
  convertSimpleData,
  SimpleReqDataItem
} from '../common/utils/api'

export const categories = writable<Category[]>([])
export const categoriesTotal = writable<number>(0)
export const categoriesStatus = writable<LoadingStatus>(LoadingStatus.None)

export const simpleCategories = writable<SimpleDataItem[]>([])
export const simpleCategoriesStatus = writable<LoadingStatus>(LoadingStatus.None)

const categoryPathDelimeter = ' | '

export const loadCategoriesTree = async () => {
  categoriesStatus.set(LoadingStatus.Loading)
  getCategoriesTree()
    .then((res: SuccessResponse) => {
      const data = res.data as Category[]
      categories.set(data)
      categoriesStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      categoriesStatus.set(LoadingStatus.Error)
    })
}

export const loadSimpleCategoriesList = async () => {
  simpleCategoriesStatus.set(LoadingStatus.Loading)
  getCategoriesList()
    .then((res: SuccessResponse) => {
      const data = convertSimpleData(res.data as SimpleReqDataItem[])
      simpleCategories.set(data)
      simpleCategoriesStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      simpleCategoriesStatus.set(LoadingStatus.Error)
    })
}

export const addCategoryToStore = async (newCategory: Category) => {
  categoriesStatus.set(LoadingStatus.Loading)
  await addNewCategory(newCategory)
    .then((res: SuccessResponse) => {
      newCategory.id = Number(res.message)

      categories.update((categories) => {
        if (!newCategory.parentId) {
          return [...categories, newCategory]
        }
        const parentCategory = findCategoryById(newCategory.parentId, categories)

        if (parentCategory) {
          parentCategory.categories = parentCategory.categories ? [...parentCategory.categories, newCategory] : [newCategory]
        } else {
          categories = [...categories, newCategory]
        }

        return categories
      })

      categoriesStatus.set(LoadingStatus.Finished)
    })
    .catch((err: ErrorResponse) => {
      console.log(err)
      categoriesStatus.set(LoadingStatus.Error)
    })

  await loadSimpleCategoriesList()
}

export const buildCategoryList = (categoryId: number | null): string[] => {
  if (!categoryId) {
    return []
  }

  const path = findParentCategory(categoryId, get(categories))

  return path.split(categoryPathDelimeter)
}

export const findParentCategory = (categoryId: number, categories: Category[]) => {
  let name = ''

  for (let i = 0; i < categories.length; i++) {
    if (categories[i].id === categoryId) {
      return categories[i].name
    }

    if (categories[i].categories && categories[i].categories.length > 0) {
      const foundName = findParentCategory(categoryId, categories[i].categories)
      if (foundName) {
        name = `${categories[i].name}${categoryPathDelimeter}${foundName}`
      }
    }
  }

  return name
}

const findCategoryById = (id: number, categories: Category[]): Category => {
  let category = null

  for (let i = 0; i < categories.length; i++) {
    if (categories[i].id === id) {
      return categories[i]
    }

    if (categories[i].categories && categories[i].categories.length > 0) {
      category = findCategoryById(id, categories[i].categories)
      if (category) {
        return category
      }
    }
  }

  return category
}

export const removeCategoryFromStore = (id: number): void => {
  categories.update((categories) => categories.filter((category: Category) => category.id !== id))
  categoriesTotal.update(total => total - 1)
}

export const updateCategoryInStore = (updatedCategory: Category) => {
  categories.update((categories) => {
    // let categoryFromStore: Category = categories.find((category: Category) => category.id === updatedCategory.id)
    let categoryFromStore: Category = findCategoryById(updatedCategory.id, categories)

    if (categoryFromStore.parentId === updatedCategory.parentId) {
      categoryFromStore = updatedCategory
      return categories
    }

    // const oldParentCategory = categories.find((category) => category.id === categoryFromStore.parentId);
    const oldParentCategory = findCategoryById(categoryFromStore.parentId, categories)
    // const newParentCategory = categories.find((category) => category.id === updatedCategory.parentId);
    const newParentCategory = findCategoryById(updatedCategory.parentId, categories)

    newParentCategory.categories = newParentCategory.categories ? [...newParentCategory.categories, updatedCategory] : [updatedCategory]
    if (oldParentCategory) {
      oldParentCategory.categories = oldParentCategory.categories.filter((category) => category.id !== updatedCategory.id)
    } else {
      categories = categories.filter((category) => category.id !== updatedCategory.id)
    }

    return categories
  })
}

export const getSimpleCategoryById = (id: number) => {
  return get(simpleCategories).find((category) => category.id === id)
}
