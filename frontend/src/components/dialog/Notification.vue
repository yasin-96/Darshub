<script lang="ts" setup>
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import { computed } from "@vue/reactivity";

const baseLayout = useBaseLayoutStore();
const isSnackBarEnabled = computed({
  get(): boolean {
    return baseLayout.currentStateOfInfoPanel;
  },
  set(tValue: boolean) {
    baseLayout.act_toggleInfoPanel(tValue);
  },
});
const infoMessage = computed(
  (): string => baseLayout.messageOfInfoPanel
);
</script>

<template>
  <div class="text-center">
    <v-snackbar v-model="isSnackBarEnabled" multi-line>
      {{ infoMessage }}

      <template v-slot:actions>
        <v-btn color="red" variant="text" @click="isSnackBarEnabled = false">
          Close Scliessen
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>
