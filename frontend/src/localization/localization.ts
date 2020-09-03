import {addMessages, init, getLocaleFromNavigator} from 'svelte-i18n'
import {DeepDictionary} from "svelte-i18n/types/runtime/types";

import en from './en'
import ru from './ru'

addMessages('en', en as DeepDictionary)
addMessages('ru', ru as DeepDictionary)

init({
    fallbackLocale: 'en',
    //initialLocale: getLocaleFromNavigator(),
    initialLocale: "en",
})