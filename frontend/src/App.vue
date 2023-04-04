<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import BaseLayout from "@/components/layout/base/BaseLayout.vue";
</script>

<template>
  <v-app>
    <v-main>
      <BaseLayout>
        <RouterView />
      </BaseLayout>
    </v-main>
  </v-app>
</template>
