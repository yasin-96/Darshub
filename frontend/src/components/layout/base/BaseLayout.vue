<script setup>
import { computed, reactive } from "vue";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";

import Navbar from "@/components/layout/base/Navbar.vue";
import Footer from "@/components/layout/base/Footer.vue";
import SidebarLeft from "@/components/layout/base/SidebarLeft.vue";

const baseLayout = useBaseLayoutStore();

const drawerLeftSide = computed(() => {
  return baseLayout.currentStateDrawerSidebarLeft;
});

const check = reactive(false);
</script>

<template>
  <div class="drawer">
    <input
      id="leftSideDrawer"
      type="checkbox"
      class="drawer-toggle"
      v-model="check"
    />
    <div class="drawer-content flex flex-col">
      <!-- Banner/ Navbar Top -->
      <slot name="info"> </slot>
      <Navbar> </Navbar>
      <slot></slot>
      <Footer></Footer>
    </div>
    <div class="drawer-side">
      <label for="leftSideDrawer" class="drawer-overlay"></label>
      <SidebarLeft />
    </div> <div class="drawer-side">
      <label for="leftSideDrawer" class="drawer-overlay"></label>
      <SidebarLeft />
    </div>
  </div>
</template>
