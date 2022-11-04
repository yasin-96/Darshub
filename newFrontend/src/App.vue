<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import BaseLayout from "@/components/layout/base/BaseLayout.vue";
import Footer from "@/components/layout/base/Footer.vue";
import Navbar from "@/components/layout/base/Navbar.vue";
</script>

<template>
  <BaseLayout>
    <template v-slot:main>
      <Navbar>
        <RouterView />
      </Navbar>
    </template>

    <template v-slot:footer>
      <!-- <Footer /> -->
    </template>
  </BaseLayout>
</template>
