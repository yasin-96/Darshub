<script setup lang="ts">
import { reactive, computed } from "vue";
import ImageUploader from "@/components/img/ImageUploader.vue";
import { useRouter } from "vue-router";

import {
  GenderOption,
  GenderOptionRO,
  type UserRequest,
} from "@/models/user/types";
import { useRegistrationStore } from "@/stores/session/registrationStore";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const router = useRouter();
const registryStore = useRegistrationStore();

const state = reactive({
  valid: false,
  passwordValid: false,
  // minDate: "1950-01-01",
  passwordRecheck: null,
  selectedCountry: {},
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

const last_name = computed({
  get() {
    return registryStore.user.last_name;
  },
  set(newLN: string) {
    registryStore.setLastName(newLN);
  },
});

const email = computed({
  get() {
    return registryStore.user.email;
  },
  set(newEmail: string) {
    registryStore.setEmail(newEmail);
  },
});

const mobileNumber = computed({
  get() {
    return registryStore.user.telNr;
  },
  set(newMobileNr: string) {
    registryStore.setMobileNumber(newMobileNr);
  },
});

const passwd = computed({
  get() {
    return registryStore.user.password;
  },
  set(newPasswd: string) {
    registryStore.setPasswd(newPasswd);
  },
});

const gender = computed({
  get() {
    return registryStore.user.bio;
  },
  set(newGender: string) {
    registryStore.setGender(newGender);
  },
});

// const validate = ()=> {
//     this.$refs.reg_registryForm.validate();
//     if (this.valid) {
//       this.$store.dispatch("registry/act_createNewUser");
//     }
//   }

const abbord = () => {
  registryStore.clear();
  router.push({ name: "index" });
};

const saveBirthday = (e: any) => {
  registryStore.setBirthday(new Date(e.target.value));
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
              v-model="first_name"
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
              v-model="last_name"
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
                @click="gender = s"
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
              v-model="email"
            />
          </label>
          <label class="block my-4">
            <span class="label-text">Mobile</span>
            <input
              type="text"
              placeholder="mobile"
              class="form-input rounded-lg w-full"
              v-model="mobileNumber"
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
                  v-model="passwd"
                />
                <button
                  class="btn btn-square"
                  @click="
                    state.fields.password.hide = !state.fields.password.hide
                  "
                >
                  <i
                    class="i-heroicons-eye w-6 h-6 block"
                    v-if="!state.fields.password.hide"
                  ></i>
                  <i class="i-heroicons-eye-slash w-6 h-6 block" v-else></i>
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
                  v-model="state.passwordRecheck"
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
