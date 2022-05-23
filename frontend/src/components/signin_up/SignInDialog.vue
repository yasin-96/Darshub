<template>
  <v-card shaped>
    <v-row justify="center">
      <v-dialog
        v-model="showDialog"
        persistent
        :max-width="sizes.mWidgth"
        :max-height="sizes.mHeight"
        shaped
      >
        <SignIn @click="closeDialog" closBtn />
      </v-dialog>
    </v-row>
  </v-card>
</template>

<script>
import { mapState } from "vuex";
import SignIn from "./SignIn.vue";
export default {
  name: "SignInDialog",
  components: {
    SignIn,
  },
  data: () => ({
    sizes: {
      mHeight: "600px",
      mWidgth: "500px",
    },
  }),
  computed: {
    ...mapState({
      signInDialogSettings: (state) => state.ui.dialogs.signInDialog
    }),
    showDialog: {
      get() {
        return this.signInDialogSettings.showDialog;
      },
      set(tValue) {
        this.$store.dispatch("ui/dialogs/act_toggleSignInDialog", tValue);
      },
    },
  },
  methods: {
    closeDialog() {
      this.showDialog = false;
    },
  },
};
</script>

<style SignIn></style>
