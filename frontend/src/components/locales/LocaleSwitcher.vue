<template>
  <v-menu offset-y>
    <template v-slot:activator="{ on, attrs }">
      <v-btn rounded small v-bind="attrs" v-on="on" outlined elevation-0 :class="`px-2 ${cssStyle}`">
        <v-icon> mdi-web </v-icon>
        <p class="font-weight-bold mt-4 mr-3">
          {{ $i18n.locale }}
        </p>
      </v-btn>
    </template>
    <v-list>
      <v-list-item v-for="(locale, index) in locales" :key="index" selectable link @click="changeLocale(locale.code)">
        <v-list-item-title> ({{ locale.code }}) {{ locale.name }} </v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script>
import { loadLocaleMessagesAsync, supportedLocales } from '@/plugins/localization/i18n';

export default {
  props:{
    cssStyle:{
      type: String
    }
  },
  data() {
    return {
      locales: supportedLocales,
      isLoading: false,
    };
  },
  methods: {
    changeLocale(locale) {
      this.isLoading = true;
      loadLocaleMessagesAsync(locale).then(() => {
        this.isLoading = false;
      });
    },
  },
};
</script>
