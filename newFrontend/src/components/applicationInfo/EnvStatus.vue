<template>
  <div v-show="devType">
    <!-- <v-chip v-if="devType === 'LOCAL'" color="grey" label
      >({{ $t("deploymentType.local") }})</v-chip
    >
    <v-chip v-if="devType === 'DEV'" color="pink" text-color="white" label
      >({{ $t("deploymentType.dev") }})</v-chip
    >
    <v-chip v-if="devType === 'PROD'" color="blue" text-color="white" label
      >({{ $t("deploymentType.uat") }})</v-chip
    > -->
  </div>
</template>

<script>
// import { mapState } from 'vuex';
const 
export default {
    name:"EnvStatus",
    // computed:{
    //     ...mapState({
    //         devType: (state) => state.appInfo.enviroment,
    //     })
    // }
};
</script>

<style></style>
