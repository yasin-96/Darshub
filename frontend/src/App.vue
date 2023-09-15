<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import BaseLayout from "@/components/layout/base/BaseLayout.vue";
import { useLoginStore } from "./stores/session/loginStore";

//Init Auth Client
const loginClient = useLoginStore();

if(loginClient.isStorageFilled){
 //async () => await loginClient.act_logUserIn()
}

</script>

<template>
  <v-app>
    <v-main>
      <BaseLayout>
        <RouterView />
      </BaseLayout>
    </v-main>
  </v-app>
</template>
