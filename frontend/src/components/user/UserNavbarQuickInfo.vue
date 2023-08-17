<script lang="ts" setup>
import { useLoginStore } from "@/stores/session/loginStore";
import { computed } from "vue";

const loginStore = useLoginStore();
const userInfo = computed(()=>{
    return loginStore.getUser
});

const settingsMenu = [
  {
    title: "Settings",
    href: "/user/settings",
    icon: "mdi-cog",
  },
];
</script>

<template>
  <div class="text-center">
    <v-menu open-on-hover>
      <template v-slot:activator="{ props }">
        <v-avatar
          v-if="loginStore?.isUserLoggedIn"
          class="ml-2 mr-2"
          v-bind="props"
        >
          <v-img :src="userInfo?.picture" alt="John"></v-img>
        </v-avatar>
      </template>

      <v-list>
        <v-list-item>
          <v-list-item-title>{{ userInfo?.nickname }}</v-list-item-title>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item
          v-for="(m, index) in settingsMenu"
          :key="index"
          :to="m.href"
        >
          <template v-slot:prepend>
            <v-icon :icon="m.icon"> </v-icon>
          </template>
          <v-list-item-title>{{ m.title }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>
