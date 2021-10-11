<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card outlined shaped>
          <v-card-title class="headline"> Darshub-Konto erstellen</v-card-title>
          <v-card-text>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-row>
                <v-col>
                  <v-text-field v-model="firstName" label="First Name" required></v-text-field>
                </v-col>
                <v-col>
                  <v-text-field v-model="lastName" label="Last Name" required></v-text-field>
                </v-col>
              </v-row>
              <v-text-field v-model="birthday" label="Birthday" required></v-text-field>
              <v-text-field v-model="profileImage" label="Name" required></v-text-field>
              <v-text-field v-model="email" label="E-mail" required></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn :disabled="!valid" color="success" class="mr-4" @click="validate"> Erstellen </v-btn>

            <v-btn color="error" class="mr-4" @click="reset"> Abbrechen </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "RegistrationView",
  conmputed: {
    ...mapState({
      newUser: (state) => state.registry.newUser,
    }),

    firstName: {
      get() {
        return this.newUser.firstName;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    lastName: {
      get() {
        return this.newUser.lastName;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    birthday: {
      get() {
        return this.newUser.birthday;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    profileImage: {
      get() {
        return this.newUser.profileImage;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    email: {
      get() {
        return this.newUser.email;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    telNr: {
      get() {
        return this.newUser.telNr;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    occupation: {
      get() {
        return this.newUser.occupation;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    company: {
      get() {
        return this.newUser.company;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    school: {
      get() {
        return this.newUser.school;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    subject: {
      get() {
        return this.newUser.subject;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    country: {
      get() {
        return this.newUser.country;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    bio: {
      get() {
        return this.newUser.bio;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    role: {
      get() {
        return this.newUser.role;
      },
      set(newValue) {
        this.$store.dispatch("", newValue);
      },
    },
    valid() {
      if (!this.firstName && !this.lastName && !this.email) {
        return false;
      }
      return true;
    },
  },
};
</script>

<style></style>
