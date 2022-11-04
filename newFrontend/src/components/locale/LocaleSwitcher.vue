<script setup lang="ts">
import { reactive, computed } from "vue";
import { useLanguageStore } from "@/stores/langStore";
import { useI18n } from "vue-i18n";

const languageStore = useLanguageStore();
const { t } = useI18n();
const state = reactive({
  selectedValue: "",
  selectedLocal: {},
});

const allLanguages = computed(() => {
  return languageStore.lang;
});

const localWithFlag = computed(() => {
  return languageStore.lang.find((x) => x.code === t.local)?.url;
});
const localISOCode = computed(() => {
  return languageStore.lang.find((x) => x.code === t.locale)?.code;
});

const changeLocale = (newLocal: string) => {
  // $i18n.locale = newLocal;
  state.selectedLocal = newLocal;
  languageStore.setCurrentLangCode(newLocal);
};

// watch(selectedLocal, (newVal, oldVal) => {
//   // console.log(newVal, oldVal);
//   // locale.value = newVal;
// });
</script>

<template>
  <div class="flex-2 px-2 mx-2">
    <div class="dropdown dropdown-end">
      <label tabindex="0" class="btn btn-ghost rounded-btn">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M10.5 21l5.25-11.25L21 21m-9-3h7.5M3 5.621a48.474 48.474 0 016-.371m0 0c1.12 0 2.233.038 3.334.114M9 5.25V3m3.334 2.364C11.176 10.658 7.69 15.08 3 17.502m9.334-12.138c.896.061 1.785.147 2.666.257m-4.589 8.495a18.023 18.023 0 01-3.827-5.802"
          />
        </svg>
      </label>
      <ul tabindex="0" class="menu dropdown-content p-2 shadow bg-base-100 rounded-box w-52 mt-4">
        <li v-for="item in allLanguages" :key="item.iso" @click="changeLocale(item.code)">
          <span :class="item.code == selectedLocal ? 'active' : ''">
            <div class="avatar placeholder">
              <div class="bg-neutral-focus text-neutral-content rounded-full w-8">
                <img :src="item.url" width="auto" height="auto" :alt="item.code" loading="lazy" />
              </div>
            </div>
            {{ item.text }}
          </span>
        </li>
      </ul>
    </div>
  </div>
</template>
