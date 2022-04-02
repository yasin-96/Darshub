<template>
  <v-row>
    <v-col cols="12" sm="12" md="12" lg="8" xl="6" class="mx-auto">
      <v-stepper outlined shaped elevation="3" alt-labels v-model="currentStep">
        <v-stepper-header>
          <v-stepper-step :complete="currentStep > 0" :step="1">
            Basis-Daten
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step :complete="currentStep > 1" :step="2">
            Weitere Daten
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step :complete="currentStep > 2" :step="3">
            Password
          </v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step :complete="currentStep > 3" :step="4">
            Abschluss
          </v-stepper-step>
        </v-stepper-header>

        <v-stepper-items>
          <v-stepper-content :step="1">
            <v-card elevation="0">
              <v-card-title>
                <v-container class="pa-5">
                  <v-row>
                    <v-col class="display-1">Darshub-Konto erstellen</v-col>
                  </v-row>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-form
                    ref="reg_registryForm"
                    v-model="valid"
                    lazy-validation
                  >
                    <v-row dense>
                      <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                        <v-text-field
                          outlined
                          data-test="first_name"
                          v-model="first_name"
                          label="First Name"
                          required
                          prepend-inner-icon="mdi-account-circle"
                        ></v-text-field>
                      </v-col>
                      <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                        <v-text-field
                          outlined
                          data-test="last_name"
                          v-model="last_name"
                          label="Last Name"
                          required
                        ></v-text-field>
                      </v-col>
                    </v-row>
                    <v-row dense>
                      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                        <v-menu
                          ref="ref_birthday"
                          v-model="birthdayModel"
                          :close-on-content-click="false"
                          transition="scale-transition"
                          min-width="auto"
                        >
                          <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                              outlined
                              data-test="birthday_input"
                              v-model="selectedBirthday"
                              label="Birthday"
                              prepend-inner-icon="mdi-calendar"
                              readonly
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>
                          </template>
                          <v-date-picker
                            v-model="selectedBirthday"
                            :active-picker.sync="birthdayPicker"
                             data-test="birthday"

                            :max="maxDate"
                            min=""
                            @change="saveBirthday"
                          ></v-date-picker>
                        </v-menu>
                      </v-col>
                      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                        <v-combobox
                          outlined
                          v-model="bio"
                          data-test="bio"
                          item-value="sex"
                          :items="genders"
                          prepend-inner-icon="mdi-gender-male-female"
                          label="Sex"
                          :return-object="false"
                        >
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
                      <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                        <v-combobox
                          outlined
                          v-model="selectedCountry"
                          data-test="country"
                          item-value="name"
                          :items="getCountriesWithFlags"
                          label="Country"
                          prepend-inner-icon="mdi-earth"
                        >
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
                  </v-form>
                </v-container>
              </v-card-text>
              <v-card-actions>
                <v-btn
                  text
                  color="error"
                  class="mx-auto text-right mr-2"
                  @click="reset"
                  data-test="abbort_baseData"

                >
                  Abbrechen
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn
                  data-test="next_baseData"
                  text
                  color="success"
                  class="ml-2"
                  @click="nextStep(2)"
                  :disabled="!checkForStepTwo"
                >
                  Weiter
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-content>
          <v-stepper-content :step="2">
            <v-card elevation="0">
              <v-card-title>
                <v-container class="pa-5">
                  <v-row>
                    <v-col class="display-1">Weitere Daten</v-col>
                  </v-row>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-form
                    ref="reg_registryForm"
                    v-model="valid"
                    lazy-validation
                  >
                    <v-row dense>
                      <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                        <v-text-field
                          outlined
                          data-test="email"
                          v-model="email"
                          label="E-mail"
                          required
                          prepend-inner-icon="mdi-email"
                        >
                        </v-text-field>
                      </v-col>
                      <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                        <v-text-field
                          outlined
                          v-model="telNr"
                          type="Number"
                          label="Mobile"
                          required
                          prepend-inner-icon="mdi-cellphone-basic"
                        >
                        </v-text-field>
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
                <v-btn
                  text
                  color="info"
                  class="mx-auto text-right mr-2"
                  @click="nextStep(1)"
                >
                  Zurück
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn
                  text
                  color="success"
                  class="ml-2"
                  @click="nextStep(3)"
                  :disabled="!checkEmailOrPhoneInput"
                >
                  Weiter
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-content>
          <v-stepper-content :step="3">
            <v-card elevation="0">
              <v-card-title>
                <v-container class="pa-5">
                  <v-row>
                    <v-col class="display-1">Password festlegen</v-col>
                  </v-row>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <v-form
                    ref="reg_registryPasswordForm"
                    v-model="passwordValid"
                    lazy-validation
                  >
                    <v-row dense justify="center">
                      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                        <v-text-field
                          outlined
                          v-model="password"
                          label="Password"
                          :type="displayPassword ? 'text' : 'password'"
                          required
                          prepend-inner-icon="mdi-lock-plus"
                          @click:append="displayPassword = !displayPassword"
                          :append-icon="
                            displayPassword ? 'mdi-eye' : 'mdi-eye-off'
                          "
                        ></v-text-field>
                      </v-col> </v-row
                    ><v-row dense justify="center">
                      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
                        <v-text-field
                          outlined
                          v-model="passwordRecheck"
                          label="Password Wiederholen"
                          :type="displayPassword ? 'text' : 'password'"
                          required
                          prepend-inner-icon="mdi-lock-check"
                          :append-icon="
                            displayPassword ? 'mdi-eye' : 'mdi-eye-off'
                          "
                        ></v-text-field>
                      </v-col>
                    </v-row>
                    <v-row dense justify="center">
                      <v-col dense cols="12" sm="12" md="6" lg="6" xl="6">
                        <v-checkbox
                          dense
                          v-model="displayPassword"
                          label="Password anzeigen"
                        ></v-checkbox>
                      </v-col>
                    </v-row>
                  </v-form>
                </v-container>
              </v-card-text>
              <v-card-actions>
                <v-btn
                  text
                  color="info"
                  class="mx-auto text-right mr-2"
                  @click="nextStep(1)"
                >
                  Zurück
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn
                  text
                  color="success"
                  class="ml-2"
                  @click="nextStep(4)"
                  :disabled="!passwordNotVerified"
                >
                  Weiter
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-content>
          <v-stepper-content :step="4">
            <v-card elevation="0">
              <v-card-title>
                <v-container class="pa-5">
                  <v-row>
                    <v-col class="display-1">Zusammenfassung</v-col>
                  </v-row>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-container class="pa-5">
                  <v-row>
                    <v-col cols="4" class="mx-auto" v-if="fullName">
                      <v-icon large>mdi-account-box-outline</v-icon>
                      <v-chip label>
                        <span>{{ fullName }}</span>
                      </v-chip>
                    </v-col>
                    <v-col cols="4" class="mx-auto">
                      <v-icon large>mdi-human-male-female</v-icon>
                      <v-chip label>
                        <span>{{ bio }}</span>
                      </v-chip>
                    </v-col>
                    <v-col cols="4" class="mx-auto">
                      <v-icon large>mdi-calendar</v-icon>
                      <v-chip label>
                        <span>{{ birthday }}</span>
                      </v-chip>
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col cols="4">
                      <v-icon large>mdi-email</v-icon>
                      <v-chip label>
                        <span>{{ email }}</span>
                      </v-chip>
                    </v-col>
                    <v-col cols="4">
                      <v-icon large>mdi-phone</v-icon>
                      <v-chip label>
                        <span>{{ telNr }}</span>
                      </v-chip>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>
              <v-card-actions>
                <v-btn
                  text
                  color="error"
                  class="mx-auto text-right mr-2"
                  @click="nextStep(3)"
                >
                  Zurück
                </v-btn>
                <v-spacer></v-spacer>
                <v-btn
                  text
                  :disabled="!checkInputFields"
                  color="success"
                  class="ml-2"
                  @click="validate"
                >
                  Erstellen
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-content>
        </v-stepper-items>
      </v-stepper>
    </v-col>
    <v-col cols="12" sm="12" md="12" lg="8" xl="6" class="mx-auto">
      <v-btn outlined>
        <!-- TODO Alex -->
        Data Privacy*
      </v-btn>
    </v-col>
  </v-row>
</template>

<script>
import { mapState, mapGetters } from "vuex";
import ImageUploader from "@/components/base/ImageUploader";
export default {
  name: "Registry",
  components: {
    ImageUploader,
  },
  data: () => ({
    valid: false,
    passwordValid: false,
    date: null,
    birthdayModel: false,
    birthdayPicker: null,
    minDate: "1950-01-01",
    currentStep: 1,
    passwordRecheck: null,
    displayPassword: false,
    selectedCountry: {},
    selectedBirthday: null,
  }),
  methods: {
    validate() {
      this.$refs.reg_registryForm.validate();
      if (this.valid) {
        this.$store.dispatch("registry/act_createNewUser");
      }
    },
    reset() {
      this.$refs.reg_registryForm.reset();
      this.$router.push({ name: "home" });
    },
    disableMenu() {
      this.birthdayModel = false;
    },
    saveBirthday(newDate) {
      console.log(newDate);
      this.birthday = new Date(newDate);
      this.$refs.ref_birthday.save(newDate);
    },
    nextStep(nextStep) {
      this.currentStep = nextStep;
    },
  },
  mounted() {
    this.$refs.reg_registryForm.validate();
  },
  computed: {
    ...mapState({
      newUser: (state) => state.registry.newUser,
      genders: (state) => state.core.genders,
      countries: (state) => state.externApi.countries,
    }),
    ...mapGetters("externApi", ["getCountriesWithFlags"]),

    maxDate() {
      return new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10);
    },
    password: {
      get() {
        return this.newUser.password;
      },
      set(newPassword) {
        this.$store.dispatch("registry/act_setPassword", newPassword);
      },
    },
    first_name: {
      get() {
        return this.newUser.first_name;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setFirstName", newValue);
      },
    },
    last_name: {
      get() {
        return this.newUser.last_name;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setLastName", newValue);
      },
    },

    birthday: {
      get() {
        if (!this.newUser.birthday) {
          return null;
        }
        console.log("Parsed" + Date.parse(Date(this.newUser.birthday)));
        // return Date(this.newUser.birthday);
        return this.newUser.birthday;
      },
      set(newValue) {
        // this.$store.dispatch("registry/act_setBirthday", new Date(newValue));
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
      if (this.newUser.first_name && this.newUser.last_name) {
        return `${this.newUser.first_name} ${this.newUser.last_name}`;
      }
      return null;
    },
    checkInputFields() {
      if (
        !this.first_name ||
        (!this.last_name && !this.email) ||
        !this.birthday ||
        !this.bio ||
        !this.country ||
        !this.telNr
      ) {
        return false;
      }
      return true;
    },
    checkForStepTwo() {
      if (
        !this.first_name ||
        !this.last_name ||
        !this.birthday ||
        !this.bio ||
        !this.country
      ) {
        return false;
      }
      return true;
    },
    checkEmailOrPhoneInput() {
      if (!this.email || !this.telNr) {
        return false;
      }
      return true;
    },
    passwordNotVerified() {
      if (!this.password || !this.passwordRecheck) {
        return false;
      }

      if (this.password != this.passwordRecheck) {
        return false;
      }

      return true;
    },
  },
  watch: {
    birthdayModel(event) {
      event && setTimeout(() => (this.birthdayPicker = "YEAR"));
    },
    selectedCountry() {
      this.country = this.selectedCountry.name;
    },
  },
};
</script>

<style></style>
