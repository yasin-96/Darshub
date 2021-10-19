<template>
  <v-container fill-height align="center" justify="center">
    <v-row>
      <v-col class="mx-auto" cols="12" sm="12" md="12" lg="8" xl="6">
        <v-card outlined shaped elevation="3">
          <v-card-title>
            <v-container class="pa-5">
              <v-row>
                <v-col class="display-1">Darshub-Konto erstellen</v-col>
                <v-col class="mx-auto text-right" v-if="fullName">
                  <v-chip label>
                    <span>{{ fullName }}</span>
                  </v-chip>
                  <v-icon large>mdi-account-box-outline</v-icon>
                </v-col>
              </v-row>
            </v-container>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-form ref="reg_registryForm" v-model="valid" lazy-validation>
                <v-row dense>
                  <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                    <v-text-field outlined v-model="firstName" label="First Name" required prepend-inner-icon="mdi-account-circle"></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                    <v-text-field outlined v-model="lastName" label="Last Name" required></v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col cols="12" sm="12" md="4" lg="4" xl="4">
                    <v-menu ref="ref_birthday" v-model="birthdayModel" :close-on-content-click="false" transition="scale-transition" offset-y min-width="auto">
                      <template v-slot:activator="{ on, attrs }">
                        <v-text-field outlined v-model="birthday" label="Birthday" prepend-inner-icon="mdi-calendar" readonly v-bind="attrs" v-on="on"></v-text-field>
                      </template>
                      <v-date-picker v-model="birthday" :active-picker.sync="birthdayPicker" :max="maxDate" min="" @change="saveBirthday"></v-date-picker>
                    </v-menu>
                  </v-col>
                  <v-col cols="12" sm="12" md="4" lg="4" xl="4">
                    <v-combobox outlined v-model="bio" item-value="sex" :items="genders" prepend-inner-icon="mdi-gender-male-female" label="Sex" :return-object="false">
                      <template v-slot:item="{ index, item }">
                        <v-icon :key="index" class="mr-3">
                          {{ item.icon }}
                        </v-icon>
                        <span>{{ item.sex }}</span>
                      </template>
                      <template v-slot:selection="data">
                        <span class="mr-3">{{ data.item }}</span>
                      </template>
                    </v-combobox>
                  </v-col>
                  <v-col cols="12" sm="12" md="4" lg="4" xl="4">
                    <v-combobox outlined v-model="country" item-value="name" :items="getCountriesWithFlags" label="Country" prepend-inner-icon="mdi-earth">
                      <template v-slot:item="{ index, item }">
                        <v-chip :key="index" label>
                          {{ item.name }}
                        </v-chip>
                        <v-spacer> </v-spacer>
                        <v-icon>{{ item.icon }}</v-icon>
                      </template>
                      <template v-slot:selection="data">
                        <v-chip>
                          <span class="mr-3">{{ data.item.icon }}</span>
                          <span>{{ data.item.name }}</span>
                        </v-chip>
                      </template>
                    </v-combobox>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                    <v-text-field outlined v-model="email" label="E-mail" required prepend-inner-icon="mdi-email"> </v-text-field>
                  </v-col>
                  <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                    <v-text-field outlined v-model="telNr" label="Mobile" required prepend-inner-icon="mdi-cellphone-basic"> </v-text-field>
                  </v-col>
                </v-row>
                <v-row dense>
                  <v-col>
                    <ImageUploader />
                  </v-col>
                </v-row>
                <v-row dense> </v-row>
              </v-form>
            </v-container>
          </v-card-text>
          <!-- <v-divider></v-divider> -->
          <v-card-actions>
            <v-btn text :disabled="!valid" color="success" class="ml-2" @click="validate"> Erstellen </v-btn>
            <v-spacer></v-spacer>
            <v-btn text color="error" class="mx-auto text-right mr-2" @click="reset"> Abbrechen </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState, mapGetters } from "vuex";
import ImageUploader from "@/components/base/ImageUploader";
export default {
  name: "RegistrationView",
  components: {
    ImageUploader,
  },
  data: () => ({
    valid: null,
    date: null,
    birthdayModel: false,
    birthdayPicker: null,
    minDate: "1950-01-01",
  }),
  methods: {
    validate() {},
    reset() {
      this.$refs.reg_registryForm.reset();
      this.$router.push({ name: "home" });
    },
    disableMenu() {
      this.birthdayModel = false;
    },
    saveBirthday(newDate) {
      console.log(newDate);
      this.$refs.ref_birthday.save(newDate);
    },
  },
  computed: {
    ...mapState({
      newUser: (state) => state.registry.newUser,
      genders: (state) => state.core.genders,
      countries: (state) => state.externApi.countries,
    }),
    ...mapGetters("externApi", ["getCountriesWithFlags"]),

    maxDate() {
      return new Date(Date.now() - new Date().getTimezoneOffset() * 60000).toISOString().substr(0, 10);
    },

    firstName: {
      get() {
        return this.newUser.firstName;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setFirstName", newValue);
      },
    },
    lastName: {
      get() {
        return this.newUser.lastName;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setLastName", newValue);
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
    fullName() {
      if (this.newUser.firstName && this.newUser.lastName) {
        return `${this.newUser.firstName} ${this.newUser.lastName}`;
      }
      return null;
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
