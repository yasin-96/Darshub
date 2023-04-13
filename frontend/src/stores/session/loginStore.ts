import { defineStore } from "pinia";
import { useFetch, useStorage } from "@vueuse/core";

import {
  UserRoles,
  type User,
  type UserLoginData,
  type UserRequest,
} from "@/models/user/types";
import type { UserDetails } from "@/models/user/interfaces";

const BACKEND_API = import.meta.env.VITE_APP_BACKEND_USER_SERVICE_URI;
const BACKEND_API_TIMEOUT = 60000;

export const useLoginStore = defineStore("loginStore", {
  state: (): UserDetails => ({
    user: {
      id: "",
      password: "",
      first_name: "",
      last_name: "",
      birthday: "",
      avatar: "",
      email: "",
      telNr: "",
      company: "",
      occupation: "",
      school: "",
      subject: "",
      country: "",
      isActive: true,
      role: new Array<UserRoles>(),
    },
  }),
  persist: true,
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
    userHasAuthorRights(): boolean {
      return this.user.role.includes(UserRoles.AUTHOR, 0) ? true : false;
    },
    userHasUserManagementRights(): boolean {
      return this.user.role.includes(UserRoles.AUTHOR, 0) ? true : false;
    },
    userHasCourseManagementRights(): boolean {
      return this.user.role.includes(UserRoles.AUTHOR, 0) ? true : false;
    },
    userHasAdminRights(): boolean {
      return this.user.role.includes(UserRoles.ADMIN, 0) ? true : false;
    },
  },
});
