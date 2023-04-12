<script setup lang="ts">
import { ref, computed } from "vue";
import LocaleSwitcher from "@/components/locale/LocaleSwitcher.vue";
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
  baseLayout.act_toggleLoginWindow(true);
};
const goToLogoutPage = () => {
  rr.push({ name: "logout" });
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
    <v-btn icon v-if="false">
      <v-icon>mdi-login</v-icon>
    </v-btn>
    <div v-if="!isUserLoggedIn">
      <button
        class="btn btn-outline flex bg-white px-2 py-2 transform hover:scale-105"
        @click="goToLoginPage()"
      >
        <span class="pl-2 pr-2">Login</span>
        <i
          class="i-heroicons-arrow-left-on-rectangle h-6 w-6 hover:text-blue-400"
        ></i>
      </button>
    </div>
    <div v-else>
      <button
        class="btn btn-outline flex bg-red-500 px-2 py-2 transform hover:scale-105"
        @click="goToLogoutPage()"
      >
        <span class="pl-2 pr-2">Logout</span>
        <i
          class="i-heroicons-arrow-right-on-rectangle h-6 w-6 hover:text-blue-400"
        ></i>
      </button>
    </div>
  </v-app-bar>

 <!-- <LoginModal /> -->
</template>
