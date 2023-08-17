import { defineStore } from "pinia";
import { useFetch, useStorage, useStorageAsync } from "@vueuse/core";
import { User, useAuth0, type Auth0VueClient } from "@auth0/auth0-vue";

import { useRoute, useRouter } from "vue-router";
import {
  UserRoles,
  type UserLoginData,
  type UserRequest,
} from "@/models/user/types";
import type { LoginDetails } from "@/models/user/interfaces";
import { authClient } from "@/plugins/auth0";

const STORGE_NAME = "_pdarshub_user";
const BACKEND_API = import.meta.env.VITE_APP_BACKEND_USER_SERVICE_URI;
const BACKEND_API_TIMEOUT = 60000;

const router = useRouter();

export const useLoginStore = defineStore("loginStore", {
  state: () => ({
    authDetails: useAuth0(),
    test: null,
  }),
  actions: {
    clear() {
      this.$reset();
      //useStorage(STORGE_NAME, useAuth0());
    },
    async act_logUserIn() {
      var response = await this.authDetails?.loginWithRedirect();

      console.log("resp ${response}");
    },
    async act_logUserOut() {
      await this.authDetails.logout();
      this.clear();
    },
  },
  getters: {
    getAvatar(): string {
      return this.authDetails?.user?.details?.picture || "";
    },
    isUserLoggedIn(): boolean {
      return this.authDetails?.isAuthenticated ? true : false;
    },
    getUserId(): string {
      return this.authDetails?.user?.details?.sub || "";
    },
    getUser(): User {
      return this.authDetails?.user?.details;
    },
    isUserAuthor(): boolean {
      if (!this.authDetails.user) {
        return false;
      }
      return this.authDetails?.user?.role.includes(UserRoles.AUTHOR, 0)
        ? true
        : false;
    },
    IsUserAccountManager(): boolean {
      if (!this.authDetails.user) {
        return false;
      }
      return this.authDetails?.user?.role.includes(UserRoles.USER_MANAGER, 0)
        ? true
        : false;
    },
    isUserCourseManager(): boolean {
      if (!this.authDetails.user) {
        return false;
      }
      return this.authDetails?.user?.role.includes(UserRoles.COURSE_MANAGER, 0)
        ? true
        : false;
    },
    isUserAdmin(): boolean {
      if (!this.authDetails.user) {
        return false;
      }
      return this.authDetails?.user?.role.includes(UserRoles.ADMIN, 0)
        ? true
        : false;
    },
  },
});
