import { get, writable } from 'svelte/store'
import {
  Category,
  SimpleDataItem
} from './types'
import { LoadingStatus } from '../../stores/utils'
import {
  getCategoriesTree,
  getCategoriesList,
  addNewCategory, removeCategory, updateCategory
} from './api'
import {
  ErrorResponse,
  SuccessResponse,
  convertSimpleData,
  SimpleReqDataItem
} from '../../common/utils/api'

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
    .then(() => {
      categories.set([])
    })
    .then(() => loadCategoriesTree())
    .then(() => loadSimpleCategoriesList())
    .catch((err: ErrorResponse) => {
      console.log(err)
      categoriesStatus.set(LoadingStatus.Error)
    })
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

export const removeCategoryFromStore = async (id: number) => {
  categoriesStatus.set(LoadingStatus.Loading)
  await removeCategory({ id })
    .then(() => {
      categories.set([])
    })
    .then(() => loadCategoriesTree())
    .then(() => loadSimpleCategoriesList())
    .catch((err: ErrorResponse) => {
      console.log(err)
      categoriesStatus.set(LoadingStatus.Error)
    })
}

export const updateCategoryInStore = async (updatedCategory: Category) => {
  categoriesStatus.set(LoadingStatus.Loading)
  await updateCategory(updatedCategory)
    .then(() => {
      categories.set([])
    })
    .then(() => loadCategoriesTree())
    .then(() => loadSimpleCategoriesList())
    .catch((err: ErrorResponse) => {
      console.log(err)
    })
}

export const getSimpleCategoryById = (id: number) => {
  return get(simpleCategories).find((category) => category.id === id)
}
