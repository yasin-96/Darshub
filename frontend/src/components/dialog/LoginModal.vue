<script lang="ts" setup>
import { reactive, computed } from "vue";
import type { UserLoginData } from "@/models/user/types";
import { useBaseLayoutStore } from "@/stores/layout/baseLayout";
import { useLoginStore } from "@/stores/session/loginStore";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";

const r = useRouter();
const { t } = useI18n();
const loginStore = useLoginStore();
const baseLayout = useBaseLayoutStore();
const state = reactive<UserLoginData>({
  email: "",
  password: "",
});

const resetLocalState = () => {
  state.email = "";
  state.password = "";
};

const loginIn = async () => {
  await loginStore.act_logUserIn({
    email: state.email,
    password: state.password,
  });
};

const closeWindows = () => {
  displayLoginWindow.value = false;
  resetLocalState();
};

const displayLoginWindow = computed({
  get() {
    return baseLayout.getCurrentStateOfLoginWindow;
  },
  set(tValue: boolean) {
    baseLayout.act_toggleLoginWindow(tValue);
  },
});

loginStore.$subscribe(() => {
  resetLocalState();
  displayLoginWindow.value = false;
  r.push({ name: "user-dashboard", params: { id: loginStore.getUserId } });
});
</script>

<template>
  <div
    :class="`relative z-10 ${!displayLoginWindow ? 'hidden' : ''}`"
    aria-labelledby="loginModal"
    role="dialog"
    aria-modal="true"
  >
    <div
      class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
    ></div>

    <div class="fixed inset-0 z-10 overflow-y-auto">
      <div
        class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
      >
        <div
          class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full max-w-md"
        >
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start justify-center">
              <div class="mt-3 sm:mt-5 sm:ml-4 sm:text-left">
                <h2
                  class="text-2xl font-medium text-center leading-6 text-gray-900"
                  id="loginModal"
                >
                  Login
                </h2>
                <div class="mt-2">
                  <div class="w-full justify-around">
                    <label class="block my-4">
                      <span class="label-text">Email</span>
                      <input
                        type="email"
                        placeholder="email"
                        class="form-input rounded-lg w-full"
                        v-model="state.email"
                      />
                    </label>
                    <label class="block my-4">
                      <span class="label-text">Password</span>

                      <div class="flex flex-row gap-2">
                        <input
                          type="password"
                          placeholder="Password"
                          class="form-input block rounded-lg w-full"
                          clearable
                          v-model="state.password"
                        />
                      </div>
                    </label>
                  </div>
                </div>
                <div class="mt-2">
                  <div class="w-full justify-around">
                    <div>
                      <hr />
                      <div @click="closeWindows">
                        <span>Noch Konto vorhanden?</span>
                        <router-link to="/registry">
                          <span class="pl-2 text-blue-500">Registrieren !</span>
                        </router-link>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div
            class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6 justify-between"
          >
            <button
              type="button"
              class="inline-flex w-full justify-center rounded-md border border-transparent bg-green-600 px-4 py-2 text-base font-medium text-white shadow-sm hover:bg-green-700 sm:ml-3 sm:w-auto sm:text-sm"
              @click="loginIn()"
            >
              Login
            </button>
            <button
              type="button"
              class="mt-3 inline-flex w-full justify-center rounded-md border border-gray-300 bg-red-600 px-4 py-2 text-base text-white font-medium shadow-sm hover:bg-red-700 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
              @click="closeWindows()"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
