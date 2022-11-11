import en from "../locales/en.js.js";
import de from "../locales/de.js.js";

import { createI18n } from "vue-i18n";

// Create VueI18n instance with options
export const i18nLanguages = createI18n({
  legacy: false,
  sync: true,
  locale: "en",
  fallbackLocale: "en",
  messages: {
    en,
    de,
  },
});
