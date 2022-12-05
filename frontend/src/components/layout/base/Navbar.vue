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
const drawer = ref(false);
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

const drawerState = computed({
  get() {
    return baseLayout.currentStateDrawerSidebarLeft;
  },
  set(tValue: boolean) {
    baseLayout.act_toggleSidebarLeft(tValue);
  },
});

const isUserLoggedIn = computed(() => {
  return loginStore.isUserLoggedIn;
});
</script>

<template>
  <header>
    <nav class="bg-gray-800 w-full z-40">
      <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
        <div class="relative flex h-16 items-center justify-between">
          <div class="flex flex-1 items-start">
            <div class="flex flex-shrink-0 items-center">
              <button
                v-if="!drawerState"
                class="hover:bg-gray-700 hover:text-white text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white rounded-full"
                @click="drawerState = true"
              >
                <i class="i-heroicons-bars-3 w-8 h-8 block m-2"></i>
              </button>
              <button
                v-else
                class="drawerButton hover:bg-gray-700 hover:text-white text-white focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white rounded-full"
                @click="drawerState = false"
              >
                <i class="i-heroicons-bars-3-bottom-left w-8 h-8 block m-2"></i>
              </button>
            </div>
          </div>

          <div
            class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0"
          >
            <LocaleSwitcher />
            <div v-if="!isUserLoggedIn">
              <button
                class="rounded-full flex bg-white px-2 py-2 transform hover:scale-105"
                @click="goToLoginPage()"
              >
                <span class="pl-2 pr-2">Login</span>
                <i
                  class="i-heroicons-arrow-left-on-rectangle h-6 w-6 hover:text-blue-400"
                ></i>
              </button>
              <LoginModal />
            </div>
            <div v-else>
              <button
                class="rounded-full flex bg-red-500 px-2 py-2 transform hover:scale-105"
                @click="goToLogoutPage()"
              >
                <span class="pl-2 pr-2">Logout</span>
                <i
                  class="i-heroicons-arrow-right-on-rectangle h-6 w-6 hover:text-blue-400"
                ></i>
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </header>
</template>
