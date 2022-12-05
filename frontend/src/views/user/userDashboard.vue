<script setup lang="ts">
import { useLoginStore } from "@/stores/session/loginStore";

const loginStore = useLoginStore();
</script>

<template>
  <div>
    <!-- User Details -->
    <!-- Statistics -->
    <!-- last Course -->
    <pre>
            {{ loginStore.getUser }}
        </pre
    >
  </div>
</template>

<style scoped></style>
