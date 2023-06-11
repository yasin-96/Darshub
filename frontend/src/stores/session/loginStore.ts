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
    authClient: useAuth0(),
    test: null
  }),
  actions: {
    clear() {
      this.$reset();
      //useStorage(STORGE_NAME, useAuth0());
    },
    async act_logUserIn() {
      await this.authClient?.loginWithRedirect();
    },
    act_logUserOut() {
      this.authClient.logout();
      this.clear();
    },
  },
  getters: {
    getAvatar(): string {
      return this.authClient?.user?.details?.picture || "";
    },
    isUserLoggedIn(): boolean {
      return this.authClient?.isAuthenticated ? true : false;
    },
    getUserId(): string {
      return this.authClient?.user?.details?.sub || "";
    },
    getUser(): User {
      return this.authClient?.user?.details;
    },
    userHasAuthorRights(): boolean {
      return this.authClient?.user.role.includes(UserRoles.AUTHOR, 0)
        ? true
        : false;
    },
    userHasUserManagementRights(): boolean {
      return this.authClient?.user.role.includes(UserRoles.AUTHOR, 0)
        ? true
        : false;
    },
    userHasCourseManagementRights(): boolean {
      return this.authClient?.user.role.includes(UserRoles.AUTHOR, 0)
        ? true
        : false;
    },
    userHasAdminRights(): boolean {
      return this.authClient?.user.role.includes(UserRoles.ADMIN, 0)
        ? true
        : false;
    },
  },
});
