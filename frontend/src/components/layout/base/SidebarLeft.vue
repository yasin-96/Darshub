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
      icon: "mdi-cast-education",
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
      icon: "mdi-email",
      href: "/contact",
    },
  ],
  userLinks: [
    {
      name: "sidebar.user.dashboard",
      icon: "mdi-view-dashboard",
      href: `/user/overview`,
    },
    {
      name: "sidebar.user.course",
      icon: "mdi-book",
      href: `/user/course`,
    },
    {
      name: "sidebar.user.settings",
      icon: "mdi-cog",
      href: `/user/settings`,
    },
  ],
  userManagement: [
    {
      name: "sidebar.admin.userManagement",
      icon: "mdi-account-group-outline",
      href: `/admin/user/management`,
    },
  ],
  courseManagement: [
    {
      name: "sidebar.admin.courseManagement",
      icon: "mdi-book-account",
      href: `/admin/course/management`,
    },
  ],
  adminLinks: [
    {
      name: "sidebar.admin.userManagement",
      icon: "mdi-account-group-outline",
      href: `/admin/user/management`,
    },
    {
      name: "sidebar.admin.courseManagement",
      icon: "mdi-book-account",
      href: `/admin/course/management`,
    },
  ],
});

const drawerLeftSide = computed({
  get() {
    return baseLayout.currentStateDrawerSidebarLeft;
  },
  set(val) {
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

    <v-list lines="one" v-if="loginStore.getUserId && loginStore.isUserCourseManager">
      <v-list-item
        v-for="link in state.courseManagement"
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

    <v-list lines="one" v-if="loginStore.getUserId && loginStore.IsUserAccountManager">
      <v-list-item
        v-for="link in state.userManagement"
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

    <v-list lines="one" v-if="loginStore.getUserId && loginStore.isUserAdmin">
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
