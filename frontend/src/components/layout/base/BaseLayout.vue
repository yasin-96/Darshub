<script setup>
import { computed, ref } from "vue";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";

const baseLayout = useBaseLayoutStore();

const drawerLeftSide = computed(() => {
  return baseLayout.currentStateDrawerSidebarLeft;
});
</script>

<template>
  <div
    :class="`flex h-screen ${
      !drawerLeftSide ? '' : 'bg-blend-screen bg-gray-700'
    }`"
  >
    <div class="flex-1 flex flex-col overflow-hidden">
      <slot name="info"> </slot>
      <slot name="nav"></slot>
      
      <div class="flex h-full">
        <slot name="sidebarLeft" class="mb-0"></slot>
        <main
          :class="`flex flex-col w-full overflow-x-hidden overflow-y-auto'
          }`"
        >
          <div class="mx-auto px-6 py-8">
            <slot name="main"></slot>
          </div>
        </main>
      </div>

      <slot name="footer"></slot>
    </div>
  </div>
</template>
