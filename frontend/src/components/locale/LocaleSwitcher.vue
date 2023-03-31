<script setup lang="ts">
import { reactive, computed, onMounted, ref } from "vue";
import { useLanguageStore } from "@/stores/langStore";
import { useI18n } from "vue-i18n";

import { onClickOutside } from "@vueuse/core";

//Component state
const state = reactive({
  selectedValue: "",
  selectedLocal: {},
  enableMenu: false,
});

const ref_languagePicker = ref();

const languageStore = useLanguageStore();
const { locale } = useI18n();

const allLanguages = computed(() => {
  return languageStore.lang;
});

const changeLocale = (newLocal: string) => {
  locale.value = newLocal;
  state.selectedLocal = newLocal;
  languageStore.setCurrentLangCode(newLocal);
  state.enableMenu = false;
};

onClickOutside(ref_languagePicker, (event) => {
  state.enableMenu = false;
});

onMounted(() => {
  state.selectedLocal = locale.value;
});
</script>

<template>
  <div :hidden="state.enableMenu" class="dropdown dropdown-bottom dropdown-end btn-rounded">
    <label tabindex="0" class="btn m-1">
      <i class="i-heroicons-language h-6 w-6 "></i>
    </label>
    <ul
      tabindex="0"
      class="dropdown-content menu p-2 shadow bg-base-100 rounded-box"
    >
      <li
        v-for="item in allLanguages"
        :key="item.iso"
        @click="changeLocale(item.code)"
      >
        <a :class="`${locale === item.code ? 'bg-amber-700 text-white' : ''}`">
          {{ item.text }}
          <img
            :src="item.url"
            width="auto"
            height="auto"
            :alt="item.code"
            loading="lazy"
          />
        </a>
      </li>
    </ul>
  </div>
</template>
