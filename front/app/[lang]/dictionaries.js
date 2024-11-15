import 'server-only'
 
const dictionaries = {
  en: () => import('../../locales/en/common.json').then((module) => module.default),
  fa: () => import('../../locales/fa/common.json').then((module) => module.default),
}
 
export const getDictionary = async (locale) => dictionaries[locale]()