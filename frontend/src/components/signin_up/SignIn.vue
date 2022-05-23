<template>
  <div>
    <v-card elevation="0">
      <div v-if="closBtn">
        <v-row dense class="ma-0 pa-0">
          <v-col class="text-end">
            <v-btn text small class="" color="error" @click="closeDialog"
              ><v-icon> mdi-close </v-icon></v-btn
            ></v-col
          >
        </v-row>
      </div>
      <v-card-title>
        <v-container class="text-center">
          <h1 class="display-1">Willkommen Zurück</h1>
        </v-container>
      </v-card-title>
      <v-card-text>
        <v-container fluid de>
          <v-form
            ref="reg_registryForm"
            v-model="signValidation"
            lazy-validation
          >
            <v-row dense>
              <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                <v-text-field
                  outlined
                  data-test="email"
                  v-model="email"
                  label="E-Mail"
                  required
                  prepend-inner-icon="mdi-email"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                <v-text-field
                  outlined
                  data-test="Password"
                  v-model="password"
                  label="Password"
                  prepend-inner-icon="mdi-lock"
                  type="Password"
                  required
                ></v-text-field>
              </v-col>
            </v-row>
          </v-form>
          <v-btn block large color="success" @click="logUserIn">Anmelden</v-btn>
        </v-container>
        <v-divider class="mx-4"></v-divider>
        <div class="text-center pa-3">
          <span class="caption"> Neu bei DarsHub? </span>
          <v-btn class="caption" text color="info" @click="goToRegView"
            >Registrieren</v-btn
          >
        </div>
      </v-card-text>
      <v-card-actions>
        <v-row dense>
          <v-col> </v-col>
          <v-col class="text-end"> </v-col>
        </v-row>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
import VueRouter from "vue-router";
const { isNavigationFailure, NavigationFailureType } = VueRouter;
export default {
  name: "SignIn",
  props:{
    closBtn: {
      type: Boolean
    }
  },
  data: () => ({
    email: null,
    password: null,
    signValidation: null,
  }),
  methods: {
    logUserIn() {
      if (this.signValidation) {
        this.$store.dispatch("userStore/user/act_logUserIn", {
          email: this.email,
          password: this.password,
        });
      }
    },
    clearInput() {
      this.password = null;
      this.email = null;
    },
    closeDialog() {
      this.$store.dispatch("ui/dialogs/act_toggleSignInDialog", false);
      this.clearInput();
    },
    goToRegView() {
      this.$store.dispatch("ui/dialogs/act_toggleSignInDialog", false);
      this.$router.push({ name: "Registry" }).catch((failure) => {
        if (isNavigationFailure(failure, NavigationFailureType.duplicated)) {
          return;
        }
      });
    },
  },
};
</script>

<style></style>
