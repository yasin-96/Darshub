<template>
  <div :class="padding" v-if="!!user && getFullName.length > 0">
    <v-btn small rounded outlined class="px-2">
      <v-icon>mdi-account-circle</v-icon>
      <span class="font-weight-bold">{{ getFullName }}</span>
    </v-btn>
  </div>
</template>

<script>
import { mapGetters, mapState } from "vuex";

export default {
  name: "UserInfo",
  props: {
    padding: {
      type: String,
    },
  },
  computed: {
    ...mapGetters({
      getFullName: ["userStore/user/getFullName"],
    }),
    ...mapState({
      user: (state) => state.userStore.user.info,
    }),
  },
};
</script>

<style></style>
