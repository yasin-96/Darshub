import type { UserInfo, UserLoginData, UserRequest } from "@/models/user/types";
import axios from "axios";
import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

interface UserDetails {
  user: UserRequest;
}

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
        ).post(logInData).json();
        if (error.value) {
          console.log("err", error.value);
          this.$reset();
        }

        if (data.value) {
          this.user = data.value;
        }
      }
    },
  },
  getters: {
    getAvatar(): string {
      return String(this.user.avatar);
    },
  },
});
