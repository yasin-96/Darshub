<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount, watch } from "vue";
import { useRoute } from "vue-router";
import { useI18n } from "vue-i18n";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import { onClickOutside } from "@vueuse/core";

import { Home, Photo, Bars3Icon, XCircleIcon } from "@heroicons/vue/24/outline";

const baseLayout = useBaseLayoutStore();

const nav_ref = ref();

const r = useRoute();
const local = useI18n();
const { t } = useI18n();

const generalSidebar = ref([
  {
    name: "sidebar.default.index",
    icon: "heroicons-home",
    disabled: false,
    href: "/",
  },
  {
    name: "sidebar.default.courses",
    icon: "",
    disabled: false,
    href: "/courses",
  },
  {
    name: "sidebar.default.gallery",
    icon: "Photo",
    disabled: false,
    href: "/gallery",
  },
]);

const generalFooter = ref([
  { name: "sidebar.footer.impress", icon: "", href: "/impressum" },
  { name: "sidebar.footer.team", icon: "", href: "/team" },
  { name: "sidebar.footer.contact", icon: "", href: "/contact" },
]);

const drawerLeftSide = computed({
  get() {
    return baseLayout.currentStateDrawerSidebarLeft;
  },
  set(val) {
    console.log("val", val);
    baseLayout.act_toggleSidebarLeft(val);
  },
});
onClickOutside(nav_ref, (event) => {
  baseLayout.act_toggleSidebarLeft(false);
});
</script>

<template>
  <!-- <transition-group name="slide-fade" :duration="2000" type="animation" mode>
    
  </transition-group> -->
  <nav
    ref="nav_ref"
    :class="`top-0 left-0  bg-white z-40 flex flex-col flex-wrap w-72 h-full justify-between shadow-2xl ${
      drawerLeftSide ? '' : 'hidden'
    }`"
  >
    <div class="w-full text-center justify-start p-4">
      <ul class="">
        <li
          v-for="link in generalSidebar"
          :key="link.name"
          class="px-1 py-1 hover:bg-slate-400 rounded-lg"
        >
          <router-link :to="link.href" class="flex flex-row px-1 py-2">
            <component :is="`${link.icon}`" class="h-6 w-6 text-black" />
            <i :class="`${link.icon} h-6 w-6`" ></i>
            <span class="pl-3" @click="drawerLeftSide != false">{{
              $t(link.name)
            }}</span>
          </router-link>
        </li>
      </ul>
    </div>

    <div class="w-full text-center justify-end p-4">
      <hr class="mb-3" />
      <ul class="">
        <li
          v-for="link in generalFooter"
          :key="link.name"
          class="px-1 py-1 hover:bg-slate-400 rounded-md"
        >
          <router-link :to="link.href" class="flex flex-row px-1 py-2">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                :d="link.icon"
              />
            </svg>
            <span class="pl-3" @click="drawerLeftSide != drawerLeftSide">{{
              $t(link.name)
            }}</span>
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>
