<template>
  <div v-show="active && message">
    <v-snackbar
      :timeout="timeout"
      :value="true"
      absolute
      right
      shaped
      top
      :color="color||defaultColor"
    >
      {{ message }}
    </v-snackbar>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "NotificationPanel",
  computed: {
    ...mapState({
      active: (state) => state.ui.notification.active,
      timeout: (state) => state.ui.notification.timeout,
      message: (state) => state.ui.notification.message,
      color: (state) => state.ui.notification.color,
      defaultColor: (state) => state.ui.notification.color,
    }),
  },
};
</script>

<style></style>
