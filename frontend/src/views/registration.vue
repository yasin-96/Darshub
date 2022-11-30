<script setup lang="ts">
import { reactive, computed } from "vue";
import { EyeSlashIcon, EyeIcon } from "@heroicons/vue/24/outline";
import ImageUploader from "@/components/img/ImageUploader.vue";
import { useRouter } from "vue-router";

import {
  GenderOption,
  GenderOptionRO,
  type UserRequest,
} from "@/models/user/types";
import { useUserRegistrationStore } from "@/stores/userRegistrationStore";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const router = useRouter();
const registryStore = useUserRegistrationStore();

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
  fields: {
    password: {
      hide: true,
    },
  },
});

const first_name = computed({
  get() {
    return registryStore.user.first_name;
  },
  set(newFn: string) {
    registryStore.setFirstName(newFn);
  },
});

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

const abbord = () => {
  registryStore.clear();
  router.push({ name: "index" });
};

const saveBirthday = (e) => {
  console.log("HIER", e.target.value);
  console.log(Date.parse(e.target.value));
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

// const checkStepOne = () => {
//   if (!registryStore.$state.user.first_name || !this.last_name || !this.birthday || !this.bio || !this.country) {
//     return false;
//   }
//   return true;
// };

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

const registerNewUser = async () => {
  await registryStore.addNewUser();
};
</script>

<template>
  <div class="min-h-screen sm:w-30 md:w-96 lg:w-96">
    <div class="flex-col flex-nowrap">
      <div class="text-center flex-col">
        <h1 class="text-5xl font-bold">TODO</h1>
      </div>

      <div class="w-full shadow-xl rounded-lg px-4 my-10">
        <div class="p-3">
          <label class="block my-6">
            <span class="label-text">{{
              $t("sites.registry.steps.first.input.firstName.label")
            }}</span>
            <input
              type="text"
              :placeholder="
                $t('sites.registry.steps.first.input.firstName.placeholder')
              "
              class="form-input rounded-lg input-bordered w-full"
              :model="first_name"
            />
          </label>

          <label class="block my-4">
            <span class="label-text">{{
              $t("sites.registry.steps.first.input.lastName.label")
            }}</span>
            <input
              type="text"
              :placeholder="
                $t('sites.registry.steps.first.input.lastName.placeholder')
              "
              class="form-input rounded-lg w-full"
            />
          </label>
          <label class="block my-4">
            <span class="label-text">{{
              $t("sites.registry.steps.first.input.birthday.label")
            }}</span>
            <input
              type="date"
              class="form-input rounded-lg w-full"
              :onchange="saveBirthday"
            />
          </label>
          <label class="block my-4">
            <span class="label-text"
              >{{ $t("sites.registry.steps.first.input.gender.label") }}
            </span>
            <select class="form-select rounded-lg w-full">
              <option
                v-for="(s, index) in GenderOptionRO"
                :key="`gender-${s}-${index}`"
              >
                {{ $t(`gender.${s}`) }}
              </option>
            </select>
          </label>
        </div>
        <div class="p-2">
          <div class="flex"></div>
        </div>
      </div>

      <div class="w-full shadow-xl rounded-lg px-4 my-10">
        <div class="p-3">
          <label class="block my-6">
            <span class="label-text">Email</span>
            <input
              type="email"
              placeholder="email"
              class="form-input rounded-lg w-full"
            />
          </label>
          <label class="block my-4">
            <span class="label-text">Mobile</span>
            <input
              type="email"
              placeholder="password"
              class="form-input rounded-lg w-full"
            />
          </label>

          <ImageUploader />
        </div>
        <div class="p-2">
          <div class="flex"></div>
        </div>
      </div>

      <div class="w-full shadow-xl rounded-lg px-4 my-10">
        <div class="p-3">
          <div class="">
            <label class="block my-4">
              <div class="flex flex-row gap-2">
                <input
                  :type="`${state.fields.password.hide ? 'password' : 'text'}`"
                  placeholder="Password"
                  class="form-input block rounded-lg w-full"
                  clearable
                />
                <button
                  class="btn btn-square"
                  @click="
                    state.fields.password.hide = !state.fields.password.hide
                  "
                >
                  <EyeIcon class="w-6 h-6" v-if="!state.fields.password.hide" />
                  <EyeSlashIcon class="w-6 h-6" v-else />
                </button>
              </div>
            </label>
          </div>

          <div class="my-4">
            <label class="block my-4">
              <div class="">
                <input
                  :type="`password`"
                  placeholder="Password Wiederholen"
                  class="form-input block rounded-lg w-full"
                />
              </div>
            </label>
          </div>
        </div>
        <div class="p-2">
          <div class="flex">
            <button
              class="btn outlined rounded-xl bg-error px-4 py-1.5 hover:bg-red-500 hover:text-white"
              @click="abbord()"
            >
              Abbrechen
            </button>
            <button
              class="btn ml-auto hover:bg-green-500 px-4 py-1.5 rounded-xl hover:text-white"
              @click="registerNewUser()"
            >
              Abchließen
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style></style>
