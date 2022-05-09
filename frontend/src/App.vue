<template>
  <v-app>
    <NotificationPanel />
    <Navbar />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import NotificationPanel from "./components/base/NotificationPanel.vue";
import Navbar from "./components/layout/Topbar/Navbar.vue";
export default {
  name: "App",
  components: {
    NotificationPanel,
    Navbar
  },
  created() {
    this.$store.dispatch("externApi/act_loadAllCountries");
  },
  data: () => ({
    //
  }),
};
</script>
