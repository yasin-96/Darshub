<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router";
import { useI18n } from "vue-i18n";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import { onClickOutside } from "@vueuse/core";

const baseLayout = useBaseLayoutStore();
const r = useRoute();
const { locale, t } = useI18n();

const nav_ref = ref();
const generalSidebar = ref([
  {
    name: "sidebar.default.index",
    icon: "i-heroicons-home",
    disabled: false,
    href: "/",
  },
  {
    name: "sidebar.default.courses",
    icon: "i-heroicons-presentation-chart-line",
    disabled: false,
    href: "/course/overview",
  },
  {
    name: "sidebar.default.gallery",
    icon: "i-heroicons-photo",
    disabled: false,
    href: "/gallery",
  },
]);

const generalFooter = ref([
  {
    name: "sidebar.footer.impress",
    icon: "i-heroicons-building-office-2",
    href: "/impressum",
  },
  {
    name: "sidebar.footer.team",
    icon: "i-heroicons-user-group",
    href: "/team",
  },
  {
    name: "sidebar.footer.contact",
    icon: "i-heroicons-envelope",
    href: "/contact",
  },
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

const closeSideBar = () => {
  baseLayout.act_toggleSidebarLeft(false);
};

onClickOutside(nav_ref, (event) => {
  baseLayout.act_toggleSidebarLeft(false);
});
</script>

<template>
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
          <router-link
            :to="link.href"
            class="flex flex-row px-1 py-2"
            @click="closeSideBar"
          >
            <i :class="`${link.icon} h-6 w-6`"></i>
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
          <router-link
            :to="link.href"
            class="flex flex-row px-1 py-2"
            @click="closeSideBar"
          >
            <i :class="`${link.icon} h-6 w-6`"></i>

            <span class="pl-3" @click="drawerLeftSide != drawerLeftSide">{{
              $t(link.name)
            }}</span>
          </router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>
