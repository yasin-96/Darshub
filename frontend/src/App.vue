<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import BaseLayout from "@/components/layout/base/BaseLayout.vue";
import Footer from "@/components/layout/base/Footer.vue";
import Navbar from "@/components/layout/base/Navbar.vue";
import SidebarLeft from "@/components/layout/base/SidebarLeft.vue";
</script>

<template>
  <BaseLayout>
    <template v-slot:nav>
      <Navbar />
    </template>

    <template v-slot:sidebarLeft>
      <SidebarLeft />
    </template>

    <template v-slot:main>
      <RouterView />
    </template>

    <template v-slot:footer>
      <Footer />
    </template>
  </BaseLayout>
</template>
