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
      await this.authDetails.loginWithRedirect();
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
      return this.authDetails?.user?.appUID || "";
    },
    getUser(): User {
      return this.authDetails?.user!;
    },
    isUserAuthor(): boolean {
      if (!this.authDetails.isAuthenticated) {
        return false;
      }
      return this.authDetails?.user?.appRoles.includes(UserRoles.AUTHOR)
        ? true
        : false;
    },
    IsUserAccountManager(): boolean {
      if (!this.authDetails.isAuthenticated) {
        return false;
      }
      return this.authDetails?.user?.appRoles.includes(UserRoles.USER_MANAGER)
        ? true
        : false;
    },
    isUserCourseManager(): boolean {
      if (!this.authDetails.isAuthenticated) {
        return false;
      }
      return this.authDetails?.user?.appRoles.includes(UserRoles.COURSE_MANAGER)
        ? true
        : false;
    },
    isUserAdmin(): boolean {
      if (!this.authDetails.isAuthenticated) {
        return false;
      }
      return this.authDetails?.user?.appRoles.includes(UserRoles.ADMIN)
        ? true
        : false;
    },
  },
});
