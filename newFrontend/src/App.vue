<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Notification from "./components/dialog/Notification.vue"
import LayoutNavbar from "./components/layout/LayoutNavbar.vue"
import LayoutFooter from "./components/layout/LayoutFooter.vue"

</script>

<template>
  <LayoutNavbar>
    <!-- <Notification class="mt-20" /> -->

    <main class="mt-10 mb-auto">
      <RouterView />

    </main>
    <LayoutFooter />
  </LayoutNavbar>
</template>