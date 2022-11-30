<script setup lang="ts">
import { ref, computed } from "vue";
import LocaleSwitcher from "@/components/locale/LocaleSwitcher.vue";
import { useRoute, useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";

const baseLayout = useBaseLayoutStore();

const r = useRoute();
const rr = useRouter();
const drawer = ref(false);
const local = useI18n();
const { t } = useI18n();

const currentSideName = ref(
  r.name ? t(`sidebar.default.${String(r.name)}`) : ""
);

const drawerState = computed({
  get() {
    return baseLayout.currentStateDrawerSidebarLeft;
  },
  set(tValue: boolean) {
    baseLayout.act_toggleSidebarLeft(tValue);
  },
});

const openSideBar = () => drawerState.value != drawerState.value;

const goToLoginPage = () => {
  rr.push({ name: "registry" });
};

// watch: {
//   currentSideName() {
//     this.currentSideName = this.$t(`sidebar.default.${String(this.r.name)}`);
//   },
// },
</script>

<template>
  <header>
    <nav class="bg-gray-800 w-full">
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
            <button
              class="rounded-full flex bg-white px-2 py-2 transform hover:scale-105"
              @click="goToLoginPage()"
            >
              <span class="pl-2 pr-2">Sign Up</span>
              <i
                class="i-heroicons-arrow-right-on-rectangle h-6 w-6 hover:text-blue-400"
              ></i>
            </button>
          </div>
        </div>
      </div>

      <!-- Mobile menu, show/hide based on menu state. -->
    </nav>

    <!-- <nav>
    <div class="drawer">
      <input id="leftSidebarDrawer" v-model="drawer" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content flex flex-col">
        <-- Navbar --
        <div class="`w-full navbar bg-base-100 fixed top-0 z-50 pt-2 pb-2`">
          <div class="flex-none">
            <label for="leftSidebarDrawer" class="btn btn-square btn-ghost">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-6 h-6 stroke-current">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
              </svg>
            </label>
          </div>
          <div class="flex-1 px-2 mx-2">{{}}</div>
          <!- <div class="flex-none hidden lg:block">
          <ul class="menu menu-horizontal">
            <li v-for="(link) in generalSidebar" :key="link.name">

              <a :href="link.href">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                  stroke="currentColor" class="w-6 h-6">
                  <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />

                </svg>
                {{$t(link.name)}}
              </a>
            </li>
          </ul>
        </div> ->

          <LocaleSwitcher />
          <router-link to="/registry">
            <button class="btn gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
              Button
            </button>
          </router-link>
        </div>

      </div>

      <!- <div class="drawer-side">
        <label for="leftSidebarDrawer" class="drawer-overlay"></label>
        <ul class="menu p-4 overflow-y-auto w-80 bg-base-100">
          <li v-for="link in generalSidebar" :key="link.name">
            <router-link :to="link.href" @click="drawer = !drawer">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />
              </svg>
              {{ $t(link.name) }}
            </router-link>
          </li>
          <li class="mt-auto"><div class="divider"></div></li>
          <li v-for="link in generalFooter" :key="link.name">
            <router-link v-if="link.href" :to="link.href" @click="drawer = !drawer">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />
              </svg>
              {{ $t(link.name) }}
            </router-link>
            <a v-else :href="link.link" target="_blank">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />
              </svg>
              {{ $t(link.name) }}
            </a>
          </li>
        </ul>
      </div> ->
    </div>
  </nav> -->
  </header>
</template>
