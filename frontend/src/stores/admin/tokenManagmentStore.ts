import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";
import type { Auth0TokenDetails, TokenResponseDetails } from "@/models/admin/types";

const AUTH0_API = import.meta.env.VITE_A0_DOMAIN_URL;
const AUTH0_USER_API_CLIENT_ID = import.meta.env.VITE_A0_USER_API_CLIENT_ID;
const AUTH0_USER_API_CLIENT_SECRET = import.meta.env
  .VITE_A0_USER_API_CLIENT_SECRET;
const AUTH0_USER_API_CLIENT_AUD = import.meta.env.VITE_A0_USER_API_CLIENT_AUD;

const AUTH0_ACCOUNT_MANAGER_API_CLIENT_ID = import.meta.env
  .VITE_A0_ACCOUNT_MANAGER_API_CLIENT_ID;
const AUTH0_ACCOUNT_MANAGER_API_CLIENT_SECRET = import.meta.env
  .VITE_A0_ACCOUNT_MANAGER_API_CLIENT_SECRET;
const AUTH0_ACCOUNT_MANAGER_API_CLIENT_AUD = import.meta.env
  .VITE_A0_ACCOUNT_MANAGER_API_CLIENT_AUD;

export const useTokenManagementStore = defineStore("tokenManagementStore", {
  state: () => ({
    userApiToken: {} as TokenResponseDetails,
    accountManagerApiToken: {} as TokenResponseDetails
  }),
  actions: {
    async getManagementTokenFromUserApi() {
      let payload = {
        client_id: AUTH0_USER_API_CLIENT_ID,
        client_secret: AUTH0_USER_API_CLIENT_SECRET,
        audience: AUTH0_USER_API_CLIENT_AUD,
        grant_type: "client_credentials",
      };

      const { error, data } = useFetch<Auth0TokenDetails>(
        `${AUTH0_API}/oauth/token`
      ).post(payload).get().json();

      if (error.value) {
        console.error("TOKEN-GEN: Token for USER-API was not created: ", error);
      }

      if (data !== null && data.value !== null) {
        console.log("USER API TOKEN")
        console.log(data)
        this.userApiToken.tokenDetails = data.value;
        this.userApiToken.lastCreation = new Date().setSeconds(
          data.value.expires_in as number,
          0
        );
      }
    },

    async getManagementTokenFromAccountManagerApi() {
      let payload = {
        client_id: AUTH0_ACCOUNT_MANAGER_API_CLIENT_ID,
        client_secret: AUTH0_ACCOUNT_MANAGER_API_CLIENT_SECRET,
        audience: AUTH0_ACCOUNT_MANAGER_API_CLIENT_AUD,
        grant_type: "client_credentials",
      };

      const { error, data } = useFetch<Auth0TokenDetails>(
        `${AUTH0_API}/oauth/token`
      ).post(payload).json();

      if (error.value) {
        console.log("TOKEN-GEN: Token for USER-API was not created: ", error);
      }

      if (data !== null && data.value !== null) {
        this.accountManagerApiToken.tokenDetails =
          data.value;
        this.accountManagerApiToken.lastCreation = new Date().setSeconds(
          data.value.expires_in as number,
          0
        );
      }
    },
  },
  getters: {
    getUserApiToken(): String {
      return this.userApiToken.tokenDetails.access_token;
    },
    getUserApiExpiredTimeInTicks(): number {
      return (
        this.userApiToken.tokenDetails.expires_in +
        this.userApiToken.lastCreation
      );
    },
    getAccountManagerApiToken(): String {
      return this.accountManagerApiToken.tokenDetails.access_token;
    },
    getAccountManagerExpiredTimeInTicks(): number {
      return (
        this.accountManagerApiToken.tokenDetails.expires_in +
        this.accountManagerApiToken.lastCreation
      );
    },
  },
});
