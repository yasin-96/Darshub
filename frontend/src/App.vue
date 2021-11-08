<template>
  <v-app>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",
  created() {
    this.$store.dispatch("externApi/act_loadAllCountries");
  },
  data: () => ({
    //
  }),
};
</script>
