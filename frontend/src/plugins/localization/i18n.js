import Vue from 'vue';
import VueI18n from 'vue-i18n';

Vue.use(VueI18n);

function getLocaleSettingOfBrowser() {
  const navigatorLocale = navigator.languages !== undefined ? navigator.languages[0] : navigator.language;

  if (!navigatorLocale) {
    return undefined;
  } else {
    const c_locale = navigatorLocale.trim().split(/-|_/)[0];
    console.log('Detected default locale of browser:', c_locale);
    return c_locale;
  }
}

const startingLocale = getLocaleSettingOfBrowser();

const i18n = new VueI18n({
  locale: startingLocale || 'de', // set locale
  fallbackLocale: 'de', // set fallback locale
  messages: {}
});

export const supportedLocales = [
  { code: 'de', name: 'Deutsch' },
  { code: 'en', name: 'English' },
];

const loadedLanguages = []

export function loadLocaleMessagesAsync(locale) {
  if (loadedLanguages.length > 0 && i18n.locale === locale) {
    return Promise.resolve(locale)
  }
  // If the language was already loaded
  if (loadedLanguages.includes(locale)) {
    i18n.locale = locale
    return Promise.resolve(locale)
  }
  // If the language hasn't been loaded yet
  return import(
    /* webpackChunkName: "locale-[request]" */ `./locales/${locale}.js`
  ).then(messages => {
    i18n.setLocaleMessage(locale, messages.default)
    loadedLanguages.push(locale)
    i18n.locale = locale
    return Promise.resolve(locale)
  })
}

loadLocaleMessagesAsync(startingLocale)

export default i18n;
