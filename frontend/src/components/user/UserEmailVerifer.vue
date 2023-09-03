<script lang="ts" setup>
import { useLoginStore } from "@/stores/session/loginStore";
import { computed } from "vue";

const loginStore = useLoginStore();
const userInfo = computed(() => {
  return loginStore.getUser;
});
</script>

<template>
  <v-app-bar :elevation="2" v-if="userInfo">
    <v-alert
      v-show="userInfo.appUID && !userInfo.appIsEmailVerfiyed"
      type="error"
      title="Email not verfied!"
      text="Please check your emails or request again verification!"
    ></v-alert>
  </v-app-bar>
</template>
