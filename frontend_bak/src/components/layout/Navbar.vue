<template>
  <div>
    <v-navigation-drawer v-model="sidebarDrawer" app>
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-app-bar-title>Application</v-app-bar-title>
    </v-app-bar>
  </div>
</template>

<script>
import { computed } from "@vue/runtime-core";
import { mapState } from "vuex";
export default {
  name: "Navbar",
  setup() {
    const sidebarDrawer = computed({
      get() {
        return this.sidebarDrawer;
      },
      set() {
        this.store.dispatch("act_toggleSidebarPanel");
      },
    });
  },
//   computed: {
//     ...mapState({
//       sidebarDrawer: (state) => state.navbar.sidebarDrawer,
//     }),
//     sidebarDrawer: {
//       get() {
//         return this.sidebarDrawer;
//       },
//       set() {
//         this.store.dispatch("act_toggleSidebarPanel");
//       },
//     },
//   },
};
</script>

<style></style>
