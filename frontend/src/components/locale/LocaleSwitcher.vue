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
  <v-btn icon>
    <v-icon>mdi-translate</v-icon>

    <v-menu activator="parent">
      <v-list>
        <v-list-item
          v-for="item in allLanguages"
          :key="item.iso"
          :value="item.text"
          :isSelected="locale === item.code"
          @click="changeLocale(item.code)"
        >
          <v-list-item-title>{{ item.text }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-btn>
</template>
