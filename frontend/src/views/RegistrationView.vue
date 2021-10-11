<template>
  <v-container fill-height align="center" justify="center">
    <v-row>
      <v-col class="mx-auto" cols="12" sm="12" md="12" lg="8" xl="6">
        <v-card outlined shaped elevation="2">
          <v-card-title class="headline">
            <v-container>
              <v-row>
                <v-col>Darshub-Konto erstellen</v-col>
              </v-row>
            </v-container>
          </v-card-title>
          <v-card-text>
            <v-form ref="reg_registryForm" v-model="valid" lazy-validation>
              <v-row dense>
                <v-col>
                  <v-text-field
                    v-model="name"
                    label="Name"
                    required
                    prepend-icon="mdi-account-circle"
                  ></v-text-field>
                </v-col>
                <v-col cols="12" sm="12" md="3" lg="3" xl="3">
                  <v-menu
                    ref="ref_birthday"
                    v-model="birthdayModel"
                    :close-on-content-click="false"
                    transition="scale-transition"
                    offset-y
                    min-width="auto"
                  >
                    <template v-slot:activator="{ on, attrs }">
                      <v-text-field
                        v-model="birthday"
                        label="Birthday"
                        prepend-icon="mdi-calendar"
                        readonly
                        v-bind="attrs"
                        v-on="on"
                      ></v-text-field>
                    </template>
                    <v-date-picker
                      v-model="birthday"
                      :active-picker.sync="birthdayPicker"
                      :max="
                        new Date(
                          Date.now() - new Date().getTimezoneOffset() * 60000
                        )
                          .toISOString()
                          .substr(0, 10)
                      "
                      min="1950-01-01"
                      @change="saveBirthday"
                    ></v-date-picker>
                  </v-menu>
                </v-col>
                <v-col cols="12" sm="12" md="2" lg="2" xl="2">
                  <v-combobox
                    v-model="bio"
                    :items="genders"
                    prepend-icon="mdi-gender-male-female"
                  >
                  </v-combobox>
                </v-col>
              </v-row>
              <v-row dense>
                <v-col>
                  <v-text-field
                    v-model="email"
                    label="E-mail"
                    required
                    prepend-icon="mdi-email"
                  >
                  </v-text-field>
                </v-col>
                <v-col cols="5">
                  <v-text-field
                    v-model="telNr"
                    label="E-mail"
                    required
                    prepend-icon="mdi-cellphone-basic"
                  >
                  </v-text-field>
                </v-col>
              </v-row>
              <v-row dense>
                <v-col>
                  <v-file-input
                    show-size
                    label="Avatar"
                    @change="selectFile"
                    prepend-icon="mdi-badge-account"
                    :hint="`Max file size is 2MB`"
                  ></v-file-input>
                </v-col>
              </v-row>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn
              text
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
            >
              Erstellen
            </v-btn>

            <v-btn text color="error" class="align-right" @click="reset">
              Abbrechen
            </v-btn>
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
    birthdayModel: false,
    birthdayPicker: null,
  }),
  methods: {
    validate() {},
    reset() {
      this.$refs.reg_registryForm.reset();
    },
    disableMenu() {
      this.birthdayModel = false;
    },
    saveBirthday(newDate) {
      console.log(newDate);
      this.$refs.ref_birthday.save(newDate);
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
  watch: {
    birthdayModel(event) {
      event && setTimeout(() => (this.birthdayPicker = "YEAR"));
    },
  },
};
</script>

<style></style>
