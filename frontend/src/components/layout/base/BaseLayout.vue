<script lang="ts" setup>
import { computed } from "vue";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";

import Navbar from "@/components/layout/base/Navbar.vue";
import Notification from "@/components/dialog/Notification.vue";
import Footer from "@/components/layout/base/Footer.vue";
import SidebarLeft from "@/components/layout/base/SidebarLeft.vue";
import UserEmailVerifer from "@/components/user/UserEmailVerifer.vue";

const baseLayout = useBaseLayoutStore();

const drawerLeftSide = computed(() => {
  return baseLayout.currentStateDrawerSidebarLeft;
});
</script>

<template>
  <v-container>
    <UserEmailVerifer />
    <Navbar />
    <SidebarLeft />
    <slot>
      <!-- Rendered Page -->
    </slot>
    <Footer />
  </v-container>
</template>
