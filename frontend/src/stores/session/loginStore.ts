import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { User, UserLoginData, UserRequest } from "@/models/user/types";
import type { UserDetails } from "@/models/user/interfaces";

const BACKEND_API = import.meta.env.VITE_APP_BACKEND_URI;
const BACKEND_API_TIMEOUT = 60000;

export const useLoginStore = defineStore("loginStore", {
  state: (): UserDetails => ({
    user: {
      id: "",
      password: "",
      first_name: "",
      last_name: "",
      birthday: null,
      avatar: "",
      email: "",
      telNr: "",
      company: "",
      occupation: "",
      school: "",
      subject: "",
      country: "",
      isActive: true,
      bio: "",
      role: new Array<Uint8Array>(),
    },
  }),
  actions: {
    clear() {
      this.$reset();
    },
    async act_logUserIn(logInData: UserLoginData) {
      if (logInData.password && logInData.email) {
        const { isFetching, error, data } = await useFetch<UserRequest>(
          `${BACKEND_API}/session`
          // { timeout: BACKEND_API_TIMEOUT }
        )
          .post(logInData)
          .json();

        if (error.value) {
          //TODO: Message popup
          console.log("err", error.value);
          this.$reset();
          return;
        }

        if (data.value) {
          this.user = data.value;
        }
      }
    },
    act_logUserOut() {
      this.$reset();
    },
  },
  getters: {
    getAvatar(): string {
      return String(this.user.avatar);
    },
    isUserLoggedIn(): boolean {
      return this.user?.id ? true : false;
    },
    getUserId(): string | null {
      return this.user.id;
    },
    getUser(): UserRequest {
      return this.user;
    },
  },
});
