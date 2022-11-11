<script setup lang="ts">
import { reactive } from "vue";

// import ImageUploader from "@/components/base/ImageUploader";
import { useRoute } from "vue-router";

import { GenderOption, GenderOptionRO } from "@/types";
import { useUserRegistrationStore } from "@/stores/userRegistrationStore";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const router = useRoute();

const state = reactive({
  valid: false,
  passwordValid: false,
  date: null,
  birthdayModel: false,
  birthdayPicker: null,
  minDate: "1950-01-01",
  currentStep: 0,
  passwordRecheck: null,
  displayPassword: false,
  selectedCountry: {},
  selectedBirthday: null,
});

const registryStore = useUserRegistrationStore();

// const validate = ()=> {
//     this.$refs.reg_registryForm.validate();
//     if (this.valid) {
//       this.$store.dispatch("registry/act_createNewUser");
//     }
//   }
const reset = () => {
  // state.$refs.reg_registryForm.reset();
  registryStore.clear();
  // router.push({ name: "home" });
};

const nextStep = (nextStep: number) => {
  state.currentStep = nextStep;
};

const saveBirthday = (e) => {

  console.log("HIER", e.target.value);
  registryStore.setBirthday(e.target.value);
};

// ...mapState({
//   newUser: (state) => state.registry.newUser,
//   genders: (state) => state.core.genders,
//   countries: (state) => state.externApi.countries,
// }),
// ...mapGetters("externApi", ["getCountriesWithFlags"]),

// const maxDate = (() => {
//   return new Date(Date.now() - new Date().getTimezoneOffset() * 60000).toISOString().substr(0, 10);
// }
// const checkInputFields = (() => {
//   if (!this.first_name || (!this.last_name && !this.email) || !this.birthday || !this.bio || !this.country || !this.telNr) {
//     return false;
//   }
//   return true;
// }
// const checkForStepTwo = (() => {
//   if (!this.first_name || !this.last_name || !this.birthday || !this.bio || !this.country) {
//     return false;
//   }
//   return true;
// }
// const checkEmailOrPhoneInput = (() => {
//   if (!this.email || !this.telNr) {
//     return false;
//   }
//   return true;
// },
// const passwordNotVerified = (() => {
//   if (!this.password || !this.passwordRecheck) {
//     return false;
//   }

//   if (this.password != this.passwordRecheck) {
//     return false;
//   }

//   return true;
// },
// },
// watch: {
//   birthdayModel(event) {
//     event && setTimeout(() => (this.birthdayPicker = "YEAR"));
//   },
//   const selectedCountry = (() => {
//     this.country = this.selectedCountry.name;
//   },
// },

const registerNewUser = async () => {
  await registryStore.addNewUser();
};
</script>

<template>
  <div class="hero min-h-screen bg-base-200">
    <div class="hero-content flex-col flex-nowrap">
      <div class="text-center flex-col">
        <h1 class="text-5xl font-bold" v-if="state.currentStep === 0">BasisDaten</h1>
        <h1 class="text-5xl font-bold" v-if="state.currentStep === 1">Weitere Daten</h1>
        <h1 class="text-5xl font-bold" v-if="state.currentStep === 2">Password</h1>
        <h1 class="text-5xl font-bold" v-if="state.currentStep === 3">Abschluss</h1>
        <div class="divider"></div>
        <div class="tabs tabs-boxed">
          <a :class="`tab ${state.currentStep === 0 ? 'tab-active' : ''}`" @click="nextStep(0)"> 1 </a>
          <a :class="`tab ${state.currentStep === 1 ? 'tab-active' : ''}`" @click="nextStep(1)"> 2</a>
          <a :class="`tab ${state.currentStep === 2 ? 'tab-active' : ''}`" @click="nextStep(2)">3</a>
        </div>
      </div>

      <div class="card max-w-sm shadow-2xl bg-base-100" v-if="state.currentStep === 0">
        <div class="card-body">
          <div class="form-control">
            <label class="label">
              <span class="label-text">{{ $t("sites.registry.steps.first.input.firstName.label") }}</span>
            </label>
            <input type="text" :placeholder="$t('sites.registry.steps.first.input.firstName.placeholder')" class="input input-bordered" />
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">{{ $t("sites.registry.steps.first.input.lastName.label") }}</span>
            </label>
            <input type="text" :placeholder="$t('sites.registry.steps.first.input.lastName.placeholder')" class="input input-bordered" />
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">{{ $t("sites.registry.steps.first.input.birthday.label") }}</span>
            </label>
            <input type="date" class="input input-bordered" :onchange="saveBirthday"/>
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">{{ $t("sites.registry.steps.first.input.gender.label") }}</span>
            </label>
            <select class="select select-bordered w-full max-w-xs">
              <option v-for="(s, index) in GenderOptionRO" :key="`gender-${s}-${index}`">
                {{ $t(`gender.${s}`) }}
              </option>
            </select>
          </div>

          <div class="form-control mt-6">
            <button class="btn btn-info" @click="nextStep(1)">next</button>
          </div>
        </div>
      </div>

      <div class="card flex-shrink-0 w-full shadow-2xl bg-base-100" v-if="state.currentStep === 1">
        <div class="card-body">
          <div class="form-control">
            <label class="label">
              <span class="label-text">Email</span>
            </label>
            <input type="text" placeholder="email" class="input input-bordered" />
          </div>
          <div class="form-control">
            <label class="label">
              <span class="label-text">Mobile</span>
            </label>
            <input type="text" placeholder="password" class="input input-bordered" />
          </div>
          <div class="form-control mt-6">
            <div class="flex">
              <button class="btn gap-2 algin-left" @click="nextStep(0)">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
                Button
              </button>
              <button class="ml-auto btn gap-2" @click="nextStep(2)">
                Button
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100" v-if="state.currentStep === 2">
        <div class="card-body">
          <div class="form-control">
            <label class="label">
              <span class="label-text">Password</span>
            </label>
            <input type="password" placeholder="password" class="input input-bordered" />
          </div>

          <div class="form-control">
            <label class="label">
              <span class="label-text">Password Wiederholen</span>
            </label>
            <input type="password" placeholder="password" class="input input-bordered" />
          </div>
          <div class="form-control mt-6">
            <div class="flex">
              <button class="btn gap-2 algin-left" @click="nextStep(1)">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
                Zurück
              </button>
              <button class="ml-auto btn gap-2" @click="registerNewUser()">
                Abschließen
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style></style>
