<script setup lang="ts">
import { ref, computed } from "vue";
import LocaleSwitcher from "@/components/locale/LocaleSwitcher.vue";
import UserNavbarQuickInfo from "@/components/user/UserNavbarQuickInfo.vue";



import { useRoute, useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import LoginModal from "@/components/dialog/LoginModal.vue";
import { useLoginStore } from "@/stores/session/loginStore";

const baseLayout = useBaseLayoutStore();
const loginStore = useLoginStore();
const r = useRoute();
const rr = useRouter();
const { t, locale } = useI18n();

const currentSideName = ref(
  r.name ? t(`sidebar.default.${String(r.name)}`) : ""
);

const goToLoginPage = () => {
  loginStore.act_logUserIn();
};
const goToLogoutPage = () => {
  loginStore.act_logUserOut();
 
  //loginStore.authClient?.loginWithRedirect();

  //rr.push({ name: "logout" });
};

const isUserLoggedIn = computed(() => {
  return loginStore.isUserLoggedIn;
});

const drawerLeftSide = computed({
  get() {
    return baseLayout.currentStateDrawerSidebarLeft;
  },
  set(val) {
    console.log("val", val);
    baseLayout.act_toggleSidebarLeft(val);
  },
});
</script>

<template>
  <v-app-bar :elevation="2">
    <v-btn icon @click="drawerLeftSide = !drawerLeftSide">
      <v-icon icon="mdi-home" />
    </v-btn>
    <v-spacer></v-spacer>
    <LocaleSwitcher />
    <UserNavbarQuickInfo />
    <v-btn icon @click="goToLoginPage()" v-if="!isUserLoggedIn">
      <v-icon>mdi-login</v-icon>
    </v-btn>
    <v-btn icon @click="goToLogoutPage()" v-if="isUserLoggedIn">
      <v-icon>mdi-logout</v-icon>
    </v-btn>

  </v-app-bar>
</template>
