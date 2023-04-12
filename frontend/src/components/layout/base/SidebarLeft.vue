<script setup lang="ts">
import { computed, reactive } from "vue";
import { useRoute } from "vue-router";
import { useI18n } from "vue-i18n";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import { useLoginStore } from "@/stores/session/loginStore";

const loginStore = useLoginStore();

const baseLayout = useBaseLayoutStore();
const r = useRoute();
const { locale, t } = useI18n();

const state = reactive({
  generalSidebar: [
    {
      name: "sidebar.default.index",
      icon: "mdi-home",
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
  <v-navigation-drawer v-model="drawerLeftSide" location="left">
    <v-list lines="one">
      <v-list-item
        v-for="link in state.generalSidebar"
        :to="link.href"
        :key="link.name"
        link
      >
        <template v-slot:prepend>
          <v-icon :icon="link.icon"></v-icon>
        </template>
        <v-list-item-title v-text="$t(link.name)"></v-list-item-title>
      </v-list-item>
    </v-list>

    <v-list lines="one" v-if="loginStore.getUserId">
      <v-list-item
        v-for="link in state.userLinks"
        :key="link.name"
        link
        :to="link.href"
      >
        <template v-slot:prepend>
          <v-icon :icon="link.icon"></v-icon>
        </template>
        <v-list-item-title v-text="$t(link.name)"></v-list-item-title>
      </v-list-item>
    </v-list>

    <v-list
      lines="one"
      v-if="loginStore.getUserId && loginStore.userHasAdminRights"
    >
      <v-list-item
        v-for="link in state.adminLinks"
        :key="link.name"
        link
        :to="link.href"
      >
        <template v-slot:prepend>
          <v-icon :icon="link.icon"></v-icon>
        </template>
        <v-list-item-title v-text="$t(link.name)"></v-list-item-title>
      </v-list-item>
    </v-list>

    <v-list lines="one">
      <v-list-item
        v-for="link in state.generalFooter"
        :key="link.name"
        link
        :to="link.href"
      >
        <template v-slot:prepend>
          <v-icon :icon="link.icon"></v-icon>
        </template>
        <v-list-item-title v-text="$t(link.name)"></v-list-item-title>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
