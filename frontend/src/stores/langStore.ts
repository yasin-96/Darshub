import { defineStore } from "pinia";

export interface LanguageCodes {
  code: string;
  iso: string;
  text: string;
  flag: string;
  url: string;
}

export const useLanguageStore = defineStore("languageStore", {
  state: () => ({
    lang: [
      {
        code: "de",
        iso: "de-DE",
        text: "DE",
        flag: "/assets/img/de1x1.svg",
        url: "https://flagicons.lipis.dev/flags/4x3/de.svg",
      },
      {
        code: "en",
        iso: "en-GB",
        text: "EN",
        flag: "/assets/img/gb1x1.svg",
        url: "https://flagicons.lipis.dev/flags/4x3/gb.svg",
      },
      // {
      //   code: "hu",
      //   iso: "hu-HU",
      //   text: "HU",
      //   flag: "/assets/img/hu1x1.svg",
      //   url: "https://flagicons.lipis.dev/flags/4x3/hu.svg",
      // },
    ] as Array<LanguageCodes>,
    currentLanguageCode: "",
  }),
  actions: {
    setCurrentLangCode(isoCode: string): void {
      this.currentLanguageCode = String(isoCode).toLowerCase();
    },
  },
  getters: {
    getIsoCode(): LanguageCodes {
      return this.lang.find((x) => x.code === this.currentLanguageCode);
    },
  },
});
