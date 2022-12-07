<script setup>
import { computed } from "vue";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";

const baseLayout = useBaseLayoutStore();

const drawerLeftSide = computed(() => {
  return baseLayout.currentStateDrawerSidebarLeft;
});
</script>

<template>
  <!-- Layer for Sidebar to hide all -->
  <div
    :class="`${
      !drawerLeftSide
        ? ''
        : 'fixed inset-0 bg-warmGray-800 bg-opacity-10 transition-opacity'
    }`"
  ></div>

  <div class="flex h-screen">
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Banner/ Navbar Top -->
      <slot name="info"> </slot>
      <slot name="nav"></slot>
      
      <div class="flex h-full">
        <!-- Sidebar left / Main Content -->
        <slot name="sidebarLeft" class=""></slot>
        <div
          :class="`flex flex-col w-full overflow-x-hidden overflow-y-auto'
          }`"
        >
          <div class="mx-auto px-6 py-8">
            <slot name="main"></slot>
          </div>
        </div>
      </div>
      <slot name="footer"></slot>
    </div>
  </div>
</template>
