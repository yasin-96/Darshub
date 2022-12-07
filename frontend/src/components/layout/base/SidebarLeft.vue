<script setup lang="ts">
import { computed, reactive, ref } from "vue";
import { useRoute } from "vue-router";
import { useI18n } from "vue-i18n";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import { onClickOutside } from "@vueuse/core";
import { useLoginStore } from "@/stores/session/loginStore";

const loginStore = useLoginStore();

const baseLayout = useBaseLayoutStore();
const r = useRoute();
const { locale, t } = useI18n();

const nav_ref = ref();

const state = reactive({
  generalSidebar: [
    {
      name: "sidebar.default.index",
      icon: "i-heroicons-home",
      disabled: false,
      href: "/",
    },
    {
      name: "sidebar.default.courses",
      icon: "i-heroicons-academic-cap",
      disabled: false,
      href: "/course/overview",
    },
  ],
  generalFooter: [
    // {
    //   name: "sidebar.footer.impress",
    //   icon: "i-heroicons-building-office-2",
    //   href: "/impressum",
    // },
    // {
    //   name: "sidebar.footer.team",
    //   icon: "i-heroicons-user-group",
    //   href: "/team",
    // },
    {
      name: "sidebar.footer.contact",
      icon: "i-heroicons-envelope",
      href: "/contact",
    },
  ],
  userLinks: [
    {
      name: "sidebar.user.dashboard",
      icon: "i-heroicons-rectangle-group",
      href: `/user/overview`,
    },
    {
      name: "sidebar.user.course",
      icon: "i-heroicons-bookmark-square",
      href: `/user/course`,
    },
    {
      name: "sidebar.user.settings",
      icon: "i-heroicons-cog-6-tooth",
      href: `/user/settings`,
    },
  ],
  adminLinks: [
    {
      name: "sidebar.admin.userManagement",
      icon: "i-heroicons-user-group",
      href: `/admin/user/management`,
    },
    {
      name: "sidebar.admin.courseManagement",
      icon: "i-heroicons-presentation-chart-bar",
      href: `/admin/course/management`,
    },
  ],
});

const closeSideBar = () => {
  baseLayout.act_toggleSidebarLeft(false);
};

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
  <nav
    ref="nav_ref"
    :class="`overflow-auto absolute top-0 left-0 bg-white z-40 flex flex-col flex-nowrap w-72 xl:w-1/4 2xl:w-1/5 h-full 
    justify-between shadow-2xl
    ${drawerLeftSide ? '' : 'hidden'}`"
  >
    <div class="w-full text-center justify-start p-2">
      <ul class="">
        <li
          v-for="link in state.generalSidebar"
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
      <hr class="my-4" />

      <ul class="" v-if="loginStore.getUserId && !!loginStore.getUserId">
        <li>
          <div
            class="text-start flex my-2 bg-blue-200 rounded-lg py-3 px-3 text-lg"
          >
            <i class="i-heroicons-user block"></i>
            <span class="pl-3">{{ $t("sidebar.user.sidebarMenuTitel") }}</span>
          </div>
          <ul class="text-md">
            <li
              v-for="link in state.userLinks"
              :key="link.name"
              class="px-3 py-1 hover:bg-slate-400 rounded-lg"
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
        </li>
      </ul>
    </div>

    <div
      class="w-full text-center justify-end p-4"
      v-if="
        loginStore.getUserId &&
        !!loginStore.getUserId &&
        loginStore.userHasAdminRights
      "
    >
      <hr class="mb-3" />

      <ul class="">
        <li>
          <div
            class="text-start flex my-2 bg-amber-700 rounded-lg py-3 px-3 text-lg"
          >
            <i class="i-heroicons-user block"></i>
            <span class="pl-3">{{ $t("sidebar.admin.sidebarMenuTitel") }}</span>
          </div>
          <ul class="">
            <li
              v-for="link in state.adminLinks"
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
        </li>
      </ul>
    </div>

    <div class="w-full text-center justify-end p-4">
      <hr class="mb-3" />
      <ul class="">
        <li
          v-for="link in state.generalFooter"
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
