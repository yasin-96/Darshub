<template>
  <v-app>
    <NotificationPanel />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import NotificationPanel from "./components/base/NotificationPanel.vue";
export default {
  name: "App",
  components: {
    NotificationPanel,
  },
  created() {
    this.$store.dispatch("externApi/act_loadAllCountries");
  },
  data: () => ({
    //
  }),
};
</script>
