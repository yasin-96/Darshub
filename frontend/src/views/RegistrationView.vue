<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card outlined shaped>
          <v-card-title class="headline"> Darshub-Konto erstellen</v-card-title>
          <v-card-text>
            <v-form ref="reg_registryForm" v-model="valid" lazy-validation>
              <v-row dense>
                <v-col>
                  <v-text-field v-model="name" label="Name" required prepend-icon="mdi-account-circle"></v-text-field>
                </v-col>
                <v-col cols="3">
                  <v-menu ref="ref_birthday" v-model="birthdayModel" :close-on-content-click="false" :return-value.sync="date" transition="scale-transition" offset-y min-width="auto">
                    <template v-slot:activator="{ on, attrs }">
                      <v-text-field v-model="date" label="Birthday" prepend-icon="mdi-calendar" readonly v-bind="attrs" v-on="on"></v-text-field>
                    </template>
                    <v-date-picker v-model="birthday" no-title scrollable>
                      <v-spacer></v-spacer>
                      <v-btn text color="primary" @click="disableMenu()"> Cancel </v-btn>
                      <v-btn text color="primary" @click="saveBirthday()"> OK </v-btn>
                    </v-date-picker>
                  </v-menu>
                </v-col>
                <v-col cols="2">
                  <v-combobox :items="genders" prepend-icon="mdi-gender-male-female"> </v-combobox>
                </v-col>
              </v-row>
              <v-row dense>
                <v-col>
                  <v-text-field v-model="email" label="E-mail" required prepend-icon="mdi-email"> </v-text-field>
                </v-col>
                <v-col cols="5">
                  <v-text-field v-model="telNr" label="E-mail" required prepend-icon="mdi-cellphone-basic"> </v-text-field>
                </v-col>
              </v-row>
              <v-row dense>
                <v-col>
                  <v-file-input show-size label="Avatar" @change="selectFile" prepend-icon="mdi-badge-account" :hint="`Max file size is 2MB`"></v-file-input>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn text :disabled="!valid" color="success" class="mr-4" @click="validate"> Erstellen </v-btn>

            <v-btn text color="error" class="align-right" @click="reset"> Abbrechen </v-btn>
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
  data: () => ({
    valid: null,
    date: null,
    birthdayModel: null,
  }),
  methods: {
    validate() {},
    reset() {
      this.$refs.reg_registryForm.reset();
    },
    disableMenu() {
      this.birthdayModel = false;
    },
    saveBirthday() {
      this.$refs.menu.save(this.date);
    },
    selectFile(newAvatarFile) {
      this.avatar = newAvatarFile;
    },
  },
  computed: {
    ...mapState({
      newUser: (state) => state.registry.newUser,
      genders: (state) => state.core.genders,
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
    name: {
      get() {
        return this.newUser.name;
      },
      set(newValue) {
        this.$store.dispatch("act_setName", newValue);
      },
    },
    birthday: {
      get() {
        return this.newUser.birthday;
      },
      set(newValue) {
        console.log(newValue);
        this.$store.dispatch("registry/act_setBirthday", newValue);
      },
    },
    avatar: {
      get() {
        return this.newUser.avatar;
      },
      set(newValue) {
        this.$store.dispatch("act_setProfileImage", newValue);
      },
    },
    email: {
      get() {
        return this.newUser.email;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setEmail", newValue);
      },
    },
    telNr: {
      get() {
        return this.newUser.telNr;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setTelNr", newValue);
      },
    },
    occupation: {
      get() {
        return this.newUser.occupation;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setOccupation", newValue);
      },
    },
    company: {
      get() {
        return this.newUser.company;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setCompany", newValue);
      },
    },
    school: {
      get() {
        return this.newUser.school;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setSchool", newValue);
      },
    },
    subject: {
      get() {
        return this.newUser.subject;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setSubject", newValue);
      },
    },
    country: {
      get() {
        return this.newUser.country;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setCountry", newValue);
      },
    },
    bio: {
      get() {
        return this.newUser.bio;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setBio", newValue);
      },
    },
    // valid() {
    //   if (!this.firstName && !this.lastName && !this.email) {
    //     return false;
    //   }
    //   return true;
    // },
  },
};
</script>

<style></style>
