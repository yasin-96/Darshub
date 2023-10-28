import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { Auth0UserExtendedObject } from "@/models/user/types";
import type { Auth0UserManagementDetails } from "@/models/admin/interfaces";
import type { Auth0UserStats } from "@/models/admin/types";
import { useTokenManagementStore } from "./tokenManagmentStore";
import type { User } from "@auth0/auth0-vue";

const AUTH0_API = import.meta.env.VITE_A0_DOMAIN_URL;
const AUTH0_API_CLIENT_ID = import.meta.env.VITE_A0_USER_API_CLIENT_ID;
const AUTH0_API_CLIENT_SECRET = import.meta.env.VITE_A0_USER_API_CLIENT_SECRET;

export const useAccountManagementStore = defineStore("accountManagementStore", {
  state: () => ({
    tokenStore: useTokenManagementStore(),
    users: [] as Array<Auth0UserExtendedObject>,
    stats: {} as Auth0UserStats,
  }),
  actions: {
    async getStats() {},
    /*
    async getAllUsersFromAuth0() {
      console.log("Start fetching data for users");
      let currentDateInTicks: number = Date.parse(new Date().toString());
      if (currentDateInTicks > this.tokenStore.getAccountManagerExpiredTimeInTicks) {
        await this.tokenStore.getManagementTokenFromUserApi();
      }
        
      console.log(`${AUTH0_API}/api/v2/users`)
      const { isFetching, error, data } = await useFetch<Array<Auth0UserExtendedObject>>(`${AUTH0_API}/api/v2/users`, (this.getAccountManagerApiToken.) {
        async beforeFetch({ url, options, cancel }) {
          if (!this.tokenStore.getAccountManagerApiToken) cancel();
          console.log(tokenStore.getManagementToken);
          options.method = "GET";
          options.redirect = 'follow';
          options.headers = {
            ...options.headers,
            Accept: "application/json",
            Authorization: `Bearer ${this.tokenStore.getManagementToken}`,
          };

          return {
            options,
          };
        },
      });


      if (error.value) {
        console.log("Error: Data from auth0 was not fetched", error);
        this.$reset();
        return;
      }

      if (data.value) {
        console.log(data.value);
        this.users = data.value;
      }
    },*/

    async getAllUsersFromAuth0() {
      console.log("Start fetching data for users");
      let currentDateInTicks: number = Date.parse(new Date().toString());
      if (
        currentDateInTicks > this.tokenStore.getAccountManagerExpiredTimeInTicks
      ) {
        await this.tokenStore.getManagementTokenFromUserApi();
      }
      const amToken = this.tokenStore.getAccountManagerApiToken 
      console.log(`${AUTH0_API}/api/v2/users`);
      const { error, data } = await useFetch<Array<Auth0UserExtendedObject>>(
        `${AUTH0_API}/api/v2/users`,
        {
          async beforeFetch({ url, options, cancel }) {

            if (!amToken) cancel();

            options.method = "GET";
            options.redirect = 'follow';
            options.headers = {
              ...options.headers,
              Accept: "application/json",
              Authorization: `Bearer ${amToken}`,
            };

            return {
              options,
            };
          },
        }
      );

      if (error.value) {
        console.log("Error: Data from auth0 was not fetched", error);
        this.$reset();
        return;
      }

      if (data !== null && data.value !== null) {
        console.log(data.value);
        this.users = data.value;
      }
    },
    async updatePersonalInformation() {},
    async updateContactInformation() {},
    async updateRole() {},
    async updateAppData() {},
    async updateMetaData() {},
  },
});
