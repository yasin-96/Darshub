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
  <div class="flex-2 px-2 mx-2">
    <button
      class="transition ease-in-out hover:scale-110 px-2 py-2 rounded-full shadow-lg text-white hover:bg-white hover:bg-white hover:text-black"
      @click="state.enableMenu = state.enableMenu ? false : true"
    >
      <i class="i-heroicons-language h-6 w-6 block"></i>
    </button>
    <transition
      enter-active-class="transition ease-out duration-100"
      enter-from-class="transform opacity-0 scale-95"
      enter-to-class="transform opacity-100 scale-100"
      leave-active-class="transition ease-in duration-75"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <div ref="ref_languagePicker"
        :class="`absolute right-32 z-40 mt-1 w-16 origin-top-right divide-gray-100 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none
        ${state.enableMenu ? '' : 'hidden'}
        `"
        role="menu"
        aria-orientation="vertical"
        aria-labelledby="menu-button"
        tabindex="-1"
      >
        <div class="py-1" role="none">
          <ul class="px-2 py-2 text-lg">
            <li
              v-for="item in allLanguages"
              :key="item.iso"
              @click="changeLocale(item.code)"
              class=""
            >
              <button
                :class="`px-1 py-1 w-full hover:bg-gray-300 hover:text-black rounded-lg ${
                  locale === item.code ? 'bg-amber-700 text-white' : ''
                }`"
                @click="changeLocale(item.code)"
              >
                {{ item.text }}
                <img
                  :src="item.url"
                  width="auto"
                  height="auto"
                  :alt="item.code"
                  loading="lazy"
                />
              </button>
            </li>
          </ul>

          <!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->
        </div>
      </div>
    </transition>
  </div>
</template>
